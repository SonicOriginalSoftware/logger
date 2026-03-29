//revive:disable:package-comments
package flat

import (
	"io"
	"log/slog"
	"os"

	"git.sonicoriginal.software/logger"
	"git.sonicoriginal.software/logger/internal/handler"
	"git.sonicoriginal.software/logger/levels"
	"git.sonicoriginal.software/logger/lib"
)

// Handler implements a flat text log handler with optional color support
type Handler struct {
	handler.Base
}

// NewHandler creates a new flat text handler that reads configuration from environment.
// LOG_LEVEL environment variable controls the log level.
// NO_COLOR environment variable disables colored output.
// Source location is automatically added when LOG_LEVEL is TRACE.
func NewHandler() *Handler {
	level := lib.ParseLevelFromEnv(logger.LevelEnvName, slog.LevelInfo)
	levelVar := &slog.LevelVar{}
	levelVar.Set(level)

	opts := &slog.HandlerOptions{
		Level: levelVar,
	}

	return &Handler{
		Base: handler.Base{
			Handler:   slog.NewTextHandler(io.Discard, opts), // Only for Enabled()
			W:         os.Stdout,
			AddSource: level <= levels.Trace,
			UseColor:  lib.ShouldUseColor(),
		},
	}
}
