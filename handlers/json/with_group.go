//revive:disable:package-comments
package json

import (
	"log/slog"

	"git.sonicoriginal.software/logger/internal/handler"
)

// WithGroup returns a new Handler with the given group
// appended to the receiver's existing groups
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Base: handler.Base{
			Handler:   h.Base.Handler.WithGroup(name),
			W:         h.Base.W,
			Attrs:     h.Base.Attrs,
			AddSource: h.Base.AddSource,
			UseColor:  h.Base.UseColor,
		},
	}
}
