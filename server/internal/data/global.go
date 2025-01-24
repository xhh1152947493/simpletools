package data

import (
	jsoniter "github.com/json-iterator/go"
	"os"
	"path/filepath"
	"simpletools/internal/configs"
	"simpletools/internal/sink"
	"simpletools/lib/logger"
	"simpletools/lib/poller"
	"sync/atomic"
)

var (
	GLog            *logger.CustomLogger  // 线程安全的全局logger
	GConfig         *configs.ServerConfig // 启动配置
	GSignalSys      chan os.Signal        // 系统信号
	GSink           *sink.EventSink       // 主线程信号接收器
	GExitMode       atomic.Int32          // 进程退出模式
	GIsExiting      atomic.Bool           // 服务器是否正在退出
	GPollerLogFlush *poller.TimePoller    // 日志Flush管理
	GTimeOffsetTs   atomic.Int64          // 游戏逻辑时间偏移量
	GUser           *OnlineUserMgr        // 登录用户管理
)

func InitGlobal(useConfig string) error {
	pathWork, err := os.Getwd()
	if err != nil {
		return err
	}
	jsonStr, err := os.ReadFile(filepath.Join(pathWork, "./etc/"+useConfig))
	if err != nil {
		return err
	}
	scfg := &configs.ServerConfig{}
	if err = jsoniter.Unmarshal(jsonStr, &scfg); err != nil {
		return err
	}
	GConfig = scfg

	GLog = logger.NewCustomLogger(GConfig.Logger, GConfig.Debug)
	Log().Info().Msg("bootstrap start...")

	now := Seconds()
	GPollerLogFlush = poller.NewTimePoller(now)

	GUser = NewOnlineUserMgr()

	GSignalSys = make(chan os.Signal, 1)
	GSink = &sink.EventSink{SinkQ: make(chan interface{}, 40000)}
	return nil
}

func Log() *logger.CustomLogger {
	return GLog
}
