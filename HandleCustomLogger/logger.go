package customlogger

import (
	"io"
	"log"
)

// Custom logger function

const (
	LevelInfo  = "INFO"
	LevelError = "ERROR"
)

type CustomLogger struct {
	logger *log.Logger
}

func NewLogger(output io.Writer, prefix string, timestamp int) *CustomLogger {
	return &CustomLogger{
		logger: log.New(output, prefix, timestamp),
	}
}

func (l *CustomLogger) Info(msg string) {
	l.logger.Printf("[%s] %s", LevelInfo, msg)
}

func (l *CustomLogger) Error(msg string) {
	l.logger.Printf("[%s] %s", LevelError, msg)
}
