package logger

import "sync"

var flusher *flusherQueue

type flusherQueue struct {
	mu sync.RWMutex
	fs []*FileLogger
}

func init() {
	flusher = &flusherQueue{}
}

func Flush() {
	flusher.mu.Lock()
	defer flusher.mu.Unlock()
	for _, f := range flusher.fs {
		f.Flush()
	}
}

func addFlusher(f *FileLogger) {
	flusher.mu.Lock()
	defer flusher.mu.Unlock()
	flusher.fs = append(flusher.fs, f)
}
