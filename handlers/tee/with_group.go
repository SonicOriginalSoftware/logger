//revive:disable:package-comments
package tee

import (
	"log/slog"
)

// WithGroup returns a new tee handler with the group applied to both handlers.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		a: h.a.WithGroup(name),
		b: h.b.WithGroup(name),
	}
}
