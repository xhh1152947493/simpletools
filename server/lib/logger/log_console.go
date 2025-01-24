package logger

import (
	"github.com/rs/zerolog"
	"time"
)

type ConsoleLogger struct {
	*zerolog.ConsoleWriter
}

func NewConsoleLogger() *ConsoleLogger {
	cw := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
	})
	return &ConsoleLogger{ConsoleWriter: &cw}
}
