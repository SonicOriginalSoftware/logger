package logger

import (
	"io"
	"log/slog"
)

// NewNullLogger creates a logger that discards all output.
// Useful for testing or when a logger is required but output is not needed.
func NewNullLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
