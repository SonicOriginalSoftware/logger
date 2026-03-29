//revive:disable:package-comments
package tee

import (
	"context"
	"log/slog"
)

// Enabled returns true if either handler is enabled at the given level.
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.a.Enabled(ctx, level) || h.b.Enabled(ctx, level)
}

// Handle sends the record to both handlers.
func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	if h.a.Enabled(ctx, record.Level) {
		if err := h.a.Handle(ctx, record); err != nil {
			return err
		}
	}
	if h.b.Enabled(ctx, record.Level) {
		return h.b.Handle(ctx, record)
	}
	return nil
}
