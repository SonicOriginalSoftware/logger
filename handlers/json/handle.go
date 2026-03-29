//revive:disable:package-comments
package json

import (
	"context"
	"log/slog"
)

// Handle the record
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	return h.Base.Handler.Handle(ctx, r)
}
