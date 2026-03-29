package logger

import (
	"context"
	"log/slog"
)

// contextKey is an unexported type for context keys to prevent collisions
// with other packages that might use string-based keys.
type contextKey struct{}

var loggerKey = contextKey{}

// FromContext extracts the logger from context.
// If no logger found, returns slog.Default().
func FromContext(ctx context.Context) *slog.Logger {
	if log, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return log
	}

	return slog.Default()
}

// ContextWithLogger returns a new context with the logger attached.
func ContextWithLogger(ctx context.Context, log *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}
