package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"simpletools/internal/data"
	"simpletools/internal/defs"
	"simpletools/lib/logger"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	wgBootstrap     sync.WaitGroup
	LastChangedDate string
	pidFile         string
)

func initConst(pidFileInput string) {
	pidFile = pidFileInput
}

func savePid() error {
	f, err := os.OpenFile(pidFile, os.O_CREATE|os.O_RDWR, 0600|0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if content, _ := io.ReadAll(f); len(content) > 0 { // 进程已经启动了，不允许重复启动
		return fmt.Errorf("process already running in pid:%s", string(content))
	}

	pid := strconv.Itoa(os.Getpid())
	if _, err = f.Write([]byte(pid)); err != nil {
		return fmt.Errorf("process write pid:%s failed", pid)
	}
	data.Log().Info().Str("pid", pid).Msg("save pid file")
	return nil
}

func dropPid() {
	err := os.Remove(pidFile)
	data.Log().Info().Err(err).Msg("drop pid file")
}

func update() {
	defer func() {
		if err := recover(); err != nil {
			tips := fmt.Sprintf("main update panic!\n%v|%s", err, string(debug.Stack()))
			data.Log().Error().Msg(tips)
		}
	}()
	now := data.Seconds()
	if data.GPollerLogFlush.PassedTime(now, 30) { // 30秒写入一次日志，避免日志长时间在内存中不进入文件
		logger.Flush()
	}
}

func preExit() {
	defer func() {
		if err := recover(); err != nil {
			tips := fmt.Sprintf("main progress pre exit failed!\n%v|%s", err, string(debug.Stack()))
			data.Log().Error().Msg(tips)
		}
	}()
	dropPid()
}

func Bootstrap(pidFileInput string, r *gin.Engine) {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("Last Changed Date: ", LastChangedDate)
		os.Exit(0)
	}
	initConst(pidFileInput)

	// 进程退出前,保证数据都已落地
	defer func() {
		if err := recover(); err != nil {
			tips := fmt.Sprintf("main progress exit by panic!\n%v|%s", err, string(debug.Stack()))
			data.Log().Error().Msg(tips)
		}
		preExit()
		data.Log().Info().Msg("main progress exit...")
		logger.Flush() // 最后调用保证所有日志全部写入
	}()

	wgBootstrap.Add(2)

	go func() { // 开启http服务器
		data.Log().Info().Msgf("http server listen on addr: %s", data.GConfig.Host)
		if !data.GConfig.Debug {
			gin.SetMode(gin.ReleaseMode)
		}
		wgBootstrap.Done()
		if err := r.Run(data.GConfig.Host); err != nil { // http server 停止后关闭整个进程
			data.Log().Error().Err(err).Msg("http server stopped")
			data.GExitMode.Store(int32(defs.ExitModePanicNoWait))
			close(data.GSink.SinkQ)
		}
	}()

	go func() { // 监听系统信号
		signal.Notify(data.GSignalSys)
		data.Log().Info().Msg("start listen system signal")
		wgBootstrap.Done()
		for {
			s := <-data.GSignalSys
			if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGKILL {
				data.Log().Info().Str("signal", s.String()).Msg("recv system exit signal")
				if s == syscall.SIGKILL {
					data.GExitMode.Store(int32(defs.ExitModeKillNoWait))
				} else {
					data.GExitMode.Store(int32(defs.ExitModeSaveAndExit))
				}
				close(data.GSink.SinkQ)
				break
			}
		}
	}()

	wgBootstrap.Wait()
	if err := savePid(); err != nil {
		data.Log().Error().Err(err).Msg("bootstrap save pid failed")
		panic(err)
	}

	data.Log().Info().Bool("IsDebug", data.GConfig.Debug).Msg("bootstrap success...")
	logger.Flush() // 立即把启动日志写入文件，便于观察是否启动成功

	ticker := time.NewTicker(time.Millisecond * 20)

Loop:
	for {
		select {
		case <-ticker.C: // 定时器驱动
			update()
		case pak, ok := <-data.GSink.SinkQ: // 事件驱动
			if ok {
				data.GSink.ConsumePak(pak)
			} else { // false对应close(data.GSink.SinkQ)通道关闭
				mode := data.GExitMode.Load()
				if mode == int32(defs.ExitModeSaveAndExit) {
					data.GIsExiting.Store(true)
					data.Log().Info().Int32("mode", mode).Msgf("main progress shutdown after %d seconds", data.GConfig.ShutdownWait)
					time.Sleep(time.Second * time.Duration(data.GConfig.ShutdownWait)) // 等待已经接收到的http请求全部完成[如果已经触发的goroutine有耗时操作的话]
				} else {
					data.Log().Info().Int32("mode", mode).Msgf("main progress shutdown now")
				}
				break Loop
			}
		}
	}
}
