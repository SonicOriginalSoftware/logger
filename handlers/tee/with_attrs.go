//revive:disable:package-comments
package tee

import (
	"log/slog"
)

// WithAttrs returns a new tee handler with attributes added to both handlers.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		a: h.a.WithAttrs(attrs),
		b: h.b.WithAttrs(attrs),
	}
}
