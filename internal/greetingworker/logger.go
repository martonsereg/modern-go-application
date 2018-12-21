package greetingworker

// Logger is the fundamental interface for all log operations.
type Logger interface {
	// Info logs an info event.
	Info(msg string, fields map[string]interface{})

	// WithFields annotates a logger with some context.
	WithFields(fields map[string]interface{}) Logger
}
