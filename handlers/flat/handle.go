//revive:disable:package-comments
package flat

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"git.sonicoriginal.software/logger/colors"
	"git.sonicoriginal.software/logger/levels"
	"git.sonicoriginal.software/logger/lib"
)

// Handle processes a log record and writes it in flat text format with optional color.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	var buf strings.Builder
	buf.Grow(256)

	// Collect all attributes (handler-level + record-level + source if enabled)
	allAttrs := lib.CollectAllAttributes(h.Base.Attrs, r, h.Base.AddSource)

	if h.Base.UseColor {
		// Determine level color
		levelColor := colors.Reset
		switch {
		case r.Level <= levels.Trace:
			levelColor = colors.Dim + colors.Cyan
		case r.Level == slog.LevelDebug:
			levelColor = colors.Gray
		case r.Level == slog.LevelInfo:
			levelColor = colors.Green
		case r.Level == slog.LevelWarn:
			levelColor = colors.Yellow
		case r.Level >= slog.LevelError:
			levelColor = colors.Red
		}

		fmt.Fprintf(&buf, "%stime=%s%s %s%s=%s%s %smsg=%s%q%s",
			colors.Cyan,
			r.Time.Format("2006-01-02T15:04:05.000Z07:00"),
			colors.Reset,
			levelColor,
			"level",
			r.Level.String(),
			colors.Reset,
			colors.Bold,
			colors.Reset,
			r.Message,
			colors.Reset)
	} else {
		fmt.Fprintf(&buf, "time=%s level=%s msg=%q",
			r.Time.Format("2006-01-02T15:04:05.000Z07:00"),
			r.Level.String(),
			r.Message)
	}

	// Add attributes
	for _, attr := range allAttrs {
		value := lib.FormatAttrValue(attr)
		if h.Base.UseColor {
			fmt.Fprintf(&buf, " %s%s=%s%s",
				colors.Dim,
				attr.Key,
				colors.Reset,
				value)
		} else {
			fmt.Fprintf(&buf, " %s=%s",
				attr.Key,
				value)
		}
	}

	buf.WriteString("\n")

	_, err := io.WriteString(h.Base.W, buf.String())
	return err
}
