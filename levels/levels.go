//revive:disable:package-comments
package levels

import "log/slog"

// Trace is a log level more verbose than Debug.
// Use this level for detailed diagnostic output including source code locations.
// Trace = -8 (Debug = -4, Info = 0, Warn = 4, Error = 8)
const Trace = slog.Level(-8)
