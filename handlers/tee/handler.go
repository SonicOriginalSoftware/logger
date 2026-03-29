//revive:disable:package-comments
package tee

import (
	"log/slog"
)

// Handler fans out log records to two slog.Handler implementations.
type Handler struct {
	a slog.Handler
	b slog.Handler
}

// NewHandler creates a tee handler that sends every log record to both handlers.
func NewHandler(a, b slog.Handler) *Handler {
	return &Handler{a: a, b: b}
}
