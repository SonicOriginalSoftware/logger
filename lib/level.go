package lib

import (
	"log/slog"
	"os"
	"strings"

	"git.sonicoriginal.software/logger/levels"
)

// ParseLevelFromEnv reads the log level from an environment variable.
// Returns the parsed level, or fallback if the env var is not set or invalid.
// Valid values: "TRACE", "DEBUG", "INFO", "WARN", "ERROR" (case-insensitive)
// Note: "WARNING" is accepted as an alias for "WARN" for compatibility.
func ParseLevelFromEnv(envVar string, fallback slog.Level) slog.Level {
	levelStr := os.Getenv(envVar)
	if levelStr == "" {
		return fallback
	}

	switch strings.ToUpper(levelStr) {
	case "TRACE":
		return levels.Trace
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return fallback
	}
}
