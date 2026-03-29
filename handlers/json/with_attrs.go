//revive:disable:package-comments
package json

import (
	"log/slog"

	"git.sonicoriginal.software/logger/internal/handler"
)

// WithAttrs returns a handler with attributes included
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Base: handler.Base{
			Handler:   h.Base.Handler.WithAttrs(attrs),
			W:         h.Base.W,
			Attrs:     h.Base.Attrs,
			AddSource: h.Base.AddSource,
			UseColor:  h.Base.UseColor,
		},
	}
}
