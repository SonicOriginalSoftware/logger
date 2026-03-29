//revive:disable:package-comments
package structured

import (
	"log/slog"

	"git.sonicoriginal.software/logger/internal/handler"
)

// WithAttrs returns a handler with attributes included
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Accumulate handler-level attributes
	newAttrs := make([]slog.Attr, len(h.Base.Attrs)+len(attrs))
	copy(newAttrs, h.Base.Attrs)
	copy(newAttrs[len(h.Base.Attrs):], attrs)

	return &Handler{
		Base: handler.Base{
			Handler:   h.Base.Handler.WithAttrs(attrs),
			W:         h.Base.W,
			Attrs:     newAttrs,
			AddSource: h.Base.AddSource,
			UseColor:  h.Base.UseColor,
		},
		AttrLevelMap: h.AttrLevelMap,
		Indents:      h.Indents,
		UnknownLevel: h.UnknownLevel,
	}
}
