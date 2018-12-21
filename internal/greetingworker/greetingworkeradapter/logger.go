package greetingworkeradapter

import (
	"github.com/goph/logur"
	"github.com/sagikazarmark/modern-go-application/internal/greetingworker"
)

// Logger wraps a logur logger and exposes it under a custom interface.
type Logger struct {
	logger logur.Logger
}

// NewLogger returns a new Logger instance.
func NewLogger(logger logur.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Info logs an info event.
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.logger.Info(msg, fields)
}

// WithFields annotates a logger with some context.
func (l *Logger) WithFields(fields map[string]interface{}) greetingworker.Logger {
	return &Logger{logger: logur.WithFields(l.logger, fields)}
}

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() *Logger {
	return NewLogger(logur.NewNoop())
}
