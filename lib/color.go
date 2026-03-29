package lib

import "os"

// ShouldUseColor determines if ANSI color codes should be used in output.
// It respects the NO_COLOR environment variable (https://no-color.org/)
// and auto-detects terminal capability if NO_COLOR is not set.
func ShouldUseColor() bool {
	// Respect NO_COLOR standard - if set to any value, disable color
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	// Auto-detect: use color if stdout is a terminal
	return IsTerminal(os.Stdout)
}
