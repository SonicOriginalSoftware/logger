//revive:disable:package-comments
package handler

import (
	"context"
	"io"
	"log/slog"
)

// Base contains common fields and methods for all log handlers
type Base struct {
	slog.Handler             // Used for Enabled() delegation
	W            io.Writer   // Output writer
	Attrs        []slog.Attr // Accumulated handler-level attributes
	AddSource    bool        // Whether to add source code location
	UseColor     bool        // Whether to use ANSI color codes
}

// Enabled delegates to the embedded handler
func (b *Base) Enabled(ctx context.Context, level slog.Level) bool {
	return b.Handler.Enabled(ctx, level)
}
