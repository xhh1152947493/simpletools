package logger

import (
	"github.com/rs/zerolog"
	"io"
)

type OptionsFunc func(l zerolog.Context) zerolog.Context

type CustomLogger struct {
	Logger zerolog.Logger
	fl     *FileLogger
}

type timestampMsHook struct{}

func (ts timestampMsHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Int64("timestamp_ms", zerolog.TimestampFunc().UnixMilli())
}

func hookTimestampMs(ctx zerolog.Context) zerolog.Context {
	th := timestampMsHook{}
	return ctx.Logger().Hook(th).With()
}

func NewCustomLogger(config Config, enableConsole bool, opts ...OptionsFunc) *CustomLogger {
	var w io.Writer
	fl := NewFileLogger(config)
	if enableConsole {
		w = zerolog.MultiLevelWriter(fl, NewConsoleLogger())
	} else {
		w = fl
	}
	ctx := zerolog.New(w).With().Timestamp()
	for _, opt := range opts {
		ctx = opt(ctx)
	}
	ctx = hookTimestampMs(ctx)                                   // 最后添加时间戳毫秒
	logger := ctx.Logger().Level(zerolog.Level(config.LogLevel)) // 设置日志级别，默认为0[debug]
	addFlusher(fl)
	return &CustomLogger{Logger: logger, fl: fl}
}

func (ml *CustomLogger) Trace() *LogEvent {
	return &LogEvent{Event: ml.Logger.Trace()}
}

func (ml *CustomLogger) Debug() *LogEvent {
	return &LogEvent{Event: ml.Logger.Debug()}
}

func (ml *CustomLogger) Info() *LogEvent {
	return &LogEvent{Event: ml.Logger.Info()}
}

func (ml *CustomLogger) Warn() *LogEvent {
	return &LogEvent{Event: ml.Logger.Warn()}
}

func (ml *CustomLogger) Error() *LogEvent {
	return &LogEvent{Event: ml.Logger.Error()}
}

func (ml *CustomLogger) Fatal() *LogEvent {
	return &LogEvent{Event: ml.Logger.Fatal()}
}

func (ml *CustomLogger) Panic() *LogEvent {
	return &LogEvent{Event: ml.Logger.Panic()}
}

func (ml *CustomLogger) Log() *LogEvent {
	return &LogEvent{Event: ml.Logger.Log()}
}
