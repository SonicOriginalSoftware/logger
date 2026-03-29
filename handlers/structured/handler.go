//revive:disable:package-comments
package structured

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"git.sonicoriginal.software/logger"
	"git.sonicoriginal.software/logger/attrs"
	"git.sonicoriginal.software/logger/internal/handler"
	"git.sonicoriginal.software/logger/levels"
	"git.sonicoriginal.software/logger/lib"
)

// Handler wraps a text handler to add hierarchical formatting and optional color to log output
type Handler struct {
	handler.Base
	Indents      [5]string              // Pre-computed indent strings for levels 0-4
	UnknownLevel int                    // Level for unknown attributes (computed from AttributeLevels)
	AttrLevelMap map[string]attrs.Level // Lookup map for attribute hierarchy
}

// NewHandler creates a new structured handler that reads configuration from environment.
// LOG_LEVEL environment variable controls the log level (TRACE, DEBUG, INFO, WARN, ERROR).
// NO_COLOR environment variable disables colored output (https://no-color.org/).
// Options can be nil to use all defaults, or provide custom AttributeLevels.
func NewHandler(opts *Options) *Handler {
	// Parse log level from environment (default: INFO)
	level := lib.ParseLevelFromEnv(logger.LevelEnvName, slog.LevelInfo)

	// Create LevelVar for potential runtime adjustment
	levelVar := &slog.LevelVar{}
	levelVar.Set(level)

	handlerOpts := &slog.HandlerOptions{
		Level: levelVar,
	}

	// Use provided AttributeLevels or fall back to defaults
	var attrLevels [][]string
	if opts != nil && opts.AttributeLevels != nil {
		attrLevels = opts.AttributeLevels
	} else {
		attrLevels = DefaultAttributeLevels
	}

	// Pre-compute indent strings for performance
	var indents [5]string
	for i := range indents {
		indents[i] = strings.Repeat("  ", i)
	}

	// Unknown level is the deepest level (len of attrLevels)
	// All attributes not in the map go here and are sorted alphabetically
	unknownLevel := len(attrLevels)

	return &Handler{
		Base: handler.Base{
			Handler:   slog.NewTextHandler(io.Discard, handlerOpts), // Only for Enabled()
			W:         os.Stdout,
			AddSource: level <= levels.Trace,
			UseColor:  lib.ShouldUseColor(),
		},
		Indents:      indents,
		UnknownLevel: unknownLevel,
		AttrLevelMap: attrs.BuildLevelMap(attrLevels),
	}
}
