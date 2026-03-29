//revive:disable:package-comments
package lib

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
)

// CollectAllAttributes gathers attributes from handler, record, and optionally source location
func CollectAllAttributes(handlerAttrs []slog.Attr, r slog.Record, addSource bool) []slog.Attr {
	// Pre-allocate with estimated capacity (handler attrs + ~8 record attrs + maybe source)
	const estimatedRecordAttrs = 8
	estimatedAttrCapacity := len(handlerAttrs) + estimatedRecordAttrs
	if addSource {
		estimatedAttrCapacity++
	}
	allAttrs := make([]slog.Attr, 0, estimatedAttrCapacity)

	// Add handler-level attributes
	allAttrs = append(allAttrs, handlerAttrs...)

	// Add source location if enabled
	if addSource {
		// Extract source location from record's PC (program counter)
		// The slog package captures this automatically when AddSource is enabled
		if r.PC != 0 {
			fs := runtime.CallersFrames([]uintptr{r.PC})
			f, _ := fs.Next()
			allAttrs = append(allAttrs, slog.Any(slog.SourceKey, &slog.Source{
				Function: f.Function,
				File:     f.File,
				Line:     f.Line,
			}))
		}
	}

	// Add record-level attributes
	r.Attrs(func(a slog.Attr) bool {
		allAttrs = append(allAttrs, a)
		return true
	})

	return allAttrs
}

// FormatAttrValue converts an attribute value to a string representation.
// Special handling for source code locations; uses type-specific formatting
// to avoid reflection overhead for common types.
func FormatAttrValue(attr slog.Attr) string {
	if attr.Key == slog.SourceKey {
		// Special formatting for source code location
		if src, ok := attr.Value.Any().(*slog.Source); ok && src != nil {
			// Format as "filename:line" for cleaner output
			return fmt.Sprintf("%s:%d", filepath.Base(src.File), src.Line)
		}
		return fmt.Sprint(attr.Value.Any())
	}

	// Use type-specific formatting to avoid reflection overhead
	switch attr.Value.Kind() {
	case slog.KindString:
		return attr.Value.String()
	case slog.KindInt64:
		return strconv.FormatInt(attr.Value.Int64(), 10)
	case slog.KindUint64:
		return strconv.FormatUint(attr.Value.Uint64(), 10)
	case slog.KindFloat64:
		return strconv.FormatFloat(attr.Value.Float64(), 'f', -1, 64)
	case slog.KindBool:
		return strconv.FormatBool(attr.Value.Bool())
	case slog.KindDuration:
		return attr.Value.Duration().String()
	case slog.KindTime:
		return attr.Value.Time().Format("2006-01-02T15:04:05.000Z07:00")
	default:
		// Complex types (Group, Any, etc.) fall back to reflection
		return fmt.Sprint(attr.Value.Any())
	}
}
