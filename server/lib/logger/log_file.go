package logger

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultFilenameFormat = "2006010215"
	defaultMaxSize        = 20
	mb                    = 1024 * 1024
	errLogName            = "clean_err.log"
)

type Config struct {
	// LogPath 日志路径，为程序运行的相对路径
	LogPath string `json:"log_path" yaml:"log_path"`

	// LogLevel 日志级别，默认Log
	LogLevel int `json:"log_level" yaml:"log_level"`

	// MaxSize 单日志文件最大大小，单位为M,默认20M
	MaxSize int `json:"max_size" yaml:"max_size"`

	// HourRotate 按小时分割日志文件，指定时则以时间分割，否则以大小MaxSize分割
	HourRotate int `json:"hour_rotate" yaml:"hour_rotate"`

	// MaxHoldDays 日志检查时间，默认不检查，单位为天，大于0则检查MaxHoldDays天以前的文件
	MaxHoldDays int `json:"max_hold_days" yaml:"max_hold_days"`

	// MaxBackups 过期日志最大保存数量数量，默认全保存。检查时过期的文件超过最大值按排序后删除最早的文件，不删除压缩后的文件
	MaxBackups int `json:"max_backups" yaml:"max_backups"`

	// EnableCompress 是否启用压缩，默认不压缩，启用压缩后对过期的日志文件进行压缩
	EnableCompress bool `json:"enable_compress" yaml:"enable_compress"`
}

// FileLogger 文件日志管理器
// 按小时分割日志文件，指定时则以时间分割，否则以大小MaxSize分割
// 支持自动删除过期日志文件
// 支持自动压缩过期日志文件
type FileLogger struct {
	config Config

	f  *os.File
	w  *bufio.Writer // 日志文件写入缓冲区
	mu sync.Mutex

	logAbsolutePath string // 日志目录绝对路径
	size            int64  // 当前文件大小
	lastRotateHour  int    // 上次分割时间
}

func NewFileLogger(config Config) *FileLogger {
	if config.MaxSize <= 0 && config.HourRotate <= 0 {
		panic(fmt.Errorf("invalid config: MaxSize and HourRotate cannot be both 0"))
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Join(pwd, config.LogPath), 0755); err != nil {
		panic(err)
	}
	fl := &FileLogger{
		config:          config,
		logAbsolutePath: filepath.Join(pwd, config.LogPath),
	}
	// 开启线程处理过期日志文件的删除和压缩
	go fl.clean()
	return fl
}

func (fl *FileLogger) cleanOnce(now time.Time) {
	earliestKeepDate := now.AddDate(0, 0, -fl.config.MaxHoldDays)
	files, err := os.ReadDir(fl.logAbsolutePath)
	if err != nil {
		logErr(now, err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		path := filepath.Join(fl.logAbsolutePath, name)
		if !strings.HasSuffix(name, ".log") {
			continue
		}

		fileTime, err := fl.fileTime(name)
		if err != nil {
			logErr(now, err)
			continue
		}

		if fileTime.After(earliestKeepDate) { // 文件时间在保留期内，不需要处理
			continue
		}

		if fl.config.EnableCompress {
			if err := fl.compressFile(path); err != nil {
				logErr(now, err)
			}
		}
	}

	if fl.config.MaxBackups > 0 {
		if err := fl.removeExpired(earliestKeepDate); err != nil {
			logErr(now, err)
		}
	}
}

// clean 遍历日志目录，删除过期日志文件，压缩过期日志文件
// 每日00:00执行一次
func (fl *FileLogger) clean() {
	for {
		now := localTime()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		duration := next.Sub(now)
		time.Sleep(duration)
		fl.cleanOnce(now)
	}
}

// fileTime 从文件名中提取时间
func (fl *FileLogger) fileTime(filename string) (time.Time, error) {
	filename = strings.TrimSuffix(filename, ".gz")
	filename = strings.TrimSuffix(filename, ".log")

	if fl.config.HourRotate > 0 { // 按小时分割的格式：2006010215
		return time.ParseInLocation(defaultFilenameFormat, filename, time.Local)
	} else { // 按大小分割的格式：2006010215_1234567890
		parts := strings.Split(filename, "_")
		if len(parts) != 2 {
			return time.Time{}, fmt.Errorf("invalid filename format: %s", filename)
		}
		timestamp, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid timestamp in filename: %s", filename)
		}
		return time.Unix(0, timestamp*int64(time.Millisecond)), nil
	}
}

// removeExpired 删除过期的日志文件,不删除gz文件
func (fl *FileLogger) removeExpired(earliestKeepDate time.Time) error {
	files, err := filepath.Glob(filepath.Join(fl.logAbsolutePath, "*.log"))
	if err != nil {
		return err
	}

	if len(files) <= fl.config.MaxBackups {
		return nil
	}

	type fileInfo struct {
		path string
		time time.Time
	}

	var fileInfos []fileInfo

	for _, file := range files {
		t, err := fl.fileTime(filepath.Base(file))
		if err != nil {
			return fmt.Errorf("error parsing file time: %v", err)
		}
		if t.After(earliestKeepDate) { // 文件时间在保留期内，不需要处理
			continue
		}
		fileInfos = append(fileInfos, fileInfo{path: file, time: t})
	}

	// 按时间降序排序
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].time.After(fileInfos[j].time)
	})

	// 删除旧文件、新文件在前面
	for _, fi := range fileInfos[fl.config.MaxBackups:] {
		if err := os.Remove(fi.path); err != nil {
			return fmt.Errorf("error removing old backup: %v", err)
		}
	}

	return nil
}

// compressFile 压缩指定的文件
func (fl *FileLogger) compressFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipPath := filePath + ".gz"
	gzipFile, err := os.Create(gzipPath)
	if err != nil {
		return err
	}
	defer gzipFile.Close()

	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, file)
	if err != nil {
		return err
	}

	// 压缩成功后删除原文件、删除前要保留旧文件
	file.Close()
	return os.Remove(filePath)
}

func (fl *FileLogger) max() int64 {
	if fl.config.MaxSize == 0 {
		return int64(defaultMaxSize * mb)
	}
	return int64(fl.config.MaxSize) * int64(mb)
}

func (fl *FileLogger) needRotate(writeLen int64) bool {
	if fl.config.HourRotate <= 0 {
		if fl.size+writeLen >= fl.max() {
			return true
		}
	} else {
		hour := localTime().Hour()
		if fl.lastRotateHour == 0 {
			fl.lastRotateHour = hour
			return true
		}
		if hour-fl.lastRotateHour >= fl.config.HourRotate {
			fl.lastRotateHour = hour
			return true
		}
	}
	return false
}

func (fl *FileLogger) openExistingOrNew() error {
	now := localTime()
	var name string
	if fl.config.HourRotate > 0 {
		name = now.Format(defaultFilenameFormat)
	} else {
		ts := now.UnixMilli()
		name = fmt.Sprintf("%s_%d", now.Format(defaultFilenameFormat), ts)
	}
	filename := fmt.Sprintf("%s/%s.log", fl.logAbsolutePath, name)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600|0644)
	if err != nil {
		return err
	}
	if fl.f != nil {
		_ = fl.w.Flush()
		_ = fl.f.Close()
	}
	fl.f = file
	fl.w = bufio.NewWriter(fl.f)
	fl.size = 0
	fl.lastRotateHour = now.Hour()
	return nil
}

func (fl *FileLogger) prepare(writeLen int64) error {
	if fl.f == nil || fl.needRotate(writeLen) {
		if err := fl.openExistingOrNew(); err != nil {
			return err
		}
	}
	return nil
}

func (fl *FileLogger) Write(p []byte) (n int, err error) {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	writeLen := int64(len(p))
	if err := fl.prepare(writeLen); err != nil {
		return 0, err
	}

	n, err = fl.w.Write(p)
	fl.size += int64(n)
	return n, err
}

func (fl *FileLogger) Flush() error {
	fl.mu.Lock()
	defer fl.mu.Unlock()
	return fl.w.Flush()
}

func localTime() time.Time {
	return time.Now().In(time.Local)
}

func log(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600|0644)
	if err != nil {
		return
	}
	defer file.Close()
	_, _ = file.WriteString(content)
}

func logErr(now time.Time, err error) {
	log(errLogName, fmt.Sprintf("%s %v\n", now.Format("2006-01-02 15:04:05"), err))
}
