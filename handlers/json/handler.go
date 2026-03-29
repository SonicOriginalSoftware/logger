//revive:disable:package-comments
package json

import (
	"log/slog"
	"os"

	"git.sonicoriginal.software/logger"
	"git.sonicoriginal.software/logger/internal/handler"
	"git.sonicoriginal.software/logger/levels"
	"git.sonicoriginal.software/logger/lib"
)

// Handler implements a JSON log handler by delegating to slog.NewJSONHandler
type Handler struct {
	handler.Base
}

// NewHandler creates a new JSON handler that reads configuration from environment.
// LOG_LEVEL environment variable controls the log level.
// Source location is automatically added when LOG_LEVEL is TRACE.
// Color is disabled since we delegate to slog.NewJSONHandler (custom colored JSON formatting is out of scope).
func NewHandler() *Handler {
	level := lib.ParseLevelFromEnv(logger.LevelEnvName, slog.LevelInfo)
	levelVar := &slog.LevelVar{}
	levelVar.Set(level)

	opts := &slog.HandlerOptions{
		Level:     levelVar,
		AddSource: level <= levels.Trace,
	}

	return &Handler{
		Base: handler.Base{
			Handler:   slog.NewJSONHandler(os.Stdout, opts),
			W:         os.Stdout,
			AddSource: level <= levels.Trace,
			UseColor:  false, // slog.NewJSONHandler doesn't support color
		},
	}
}
