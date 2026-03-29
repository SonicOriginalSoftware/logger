//revive:disable:package-comments
package structured

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"strings"

	"git.sonicoriginal.software/logger/colors"
	"git.sonicoriginal.software/logger/levels"
	"git.sonicoriginal.software/logger/lib"
)

// attrWithLevel pairs an attribute with its hierarchical level and order for rendering.
// This is used internally to sort and group attributes by their configured nesting level,
// allowing attributes to be displayed at different indentation depths (e.g., service-level
// context at level 1, span context at level 2, operation details at level 3, etc.).
type attrWithLevel struct {
	attr  slog.Attr // The actual log attribute
	level int       // Hierarchical nesting level (0=inline, 1=2 spaces, 2=4 spaces, etc.)
	order int       // Display order within the level
}

// groupAttributesByLevel organizes attributes into hierarchical levels for rendering
func (h *Handler) groupAttributesByLevel(attrs []slog.Attr) map[int][]attrWithLevel {
	unknownLevel := h.UnknownLevel

	// Pre-allocate map with capacity based on configured levels
	levelGroups := make(map[int][]attrWithLevel, unknownLevel+1)

	for _, attr := range attrs {
		if levelInfo, ok := h.AttrLevelMap[attr.Key]; ok {
			// Known attribute - use configured level and order
			levelGroups[levelInfo.Level] = append(levelGroups[levelInfo.Level], attrWithLevel{
				attr:  attr,
				level: levelInfo.Level,
				order: levelInfo.Order,
			})
		} else {
			// Unknown attribute - put at deepest level, will be sorted alphabetically
			levelGroups[unknownLevel] = append(levelGroups[unknownLevel], attrWithLevel{
				attr:  attr,
				level: unknownLevel,
				order: 0, // Will be sorted by key name instead
			})
		}
	}

	// Sort attributes within each level by order
	for level := range levelGroups {
		attrs := levelGroups[level]
		if level == unknownLevel {
			// Unknown attributes: sort alphabetically by key
			sort.Slice(attrs, func(i, j int) bool {
				return attrs[i].attr.Key < attrs[j].attr.Key
			})
		} else {
			// Known attributes: sort by configured order
			sort.Slice(attrs, func(i, j int) bool {
				return attrs[i].order < attrs[j].order
			})
		}
	}

	return levelGroups
}

// renderAttributes writes formatted attributes to the buffer at their hierarchical levels
func (h *Handler) renderAttributes(buf *strings.Builder, levelGroups map[int][]attrWithLevel) {
	unknownLevel := h.UnknownLevel

	// Render each level in order (1, 2, 3, 4+)
	// Level 0 is for inline attributes (currently none)
	for level := 1; level <= unknownLevel; level++ {
		attrs, ok := levelGroups[level]
		if !ok || len(attrs) == 0 {
			continue
		}

		// Use pre-computed indent string (optimization: avoids string allocation)
		// Fallback to dynamic generation if level exceeds pre-computed indents
		var indent string
		if level < len(h.Indents) {
			indent = h.Indents[level]
		} else {
			indent = strings.Repeat("  ", level)
		}

		for _, attrInfo := range attrs {
			attr := attrInfo.attr
			value := lib.FormatAttrValue(attr)

			// Format: INDENT KEY: VALUE
			buf.WriteString(indent)
			buf.WriteString(colors.Dim)
			buf.WriteString(attr.Key)
			buf.WriteString(colors.Reset)
			buf.WriteString(": ")
			buf.WriteString(value)
			buf.WriteString("\n")
		}
	}
}

// Handle processes a log record and writes it with color formatting.
// The context parameter is required by the slog.Handler interface but is not used
// in this implementation since all necessary information is in the record.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	// Pre-allocate string builder with reasonable capacity
	var buf strings.Builder
	buf.Grow(512)

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

		fmt.Fprintf(&buf, "%s%s%s %s%s%s %s%s%s\n",
			colors.Cyan,
			r.Time.Format("2006-01-02T15:04:05.000Z07:00"),
			colors.Reset,
			levelColor,
			r.Level.String(),
			colors.Reset,
			colors.Bold,
			r.Message,
			colors.Reset)
	} else {
		fmt.Fprintf(&buf, "%s %s %s\n",
			r.Time.Format("2006-01-02T15:04:05.000Z07:00"),
			r.Level.String(),
			r.Message)
	}

	// Group and sort attributes by hierarchical level
	levelGroups := h.groupAttributesByLevel(allAttrs)

	// Render attributes at each level with proper indentation
	h.renderAttributes(&buf, levelGroups)

	// Single write to output (optimization: batches all output)
	_, err := io.WriteString(h.Base.W, buf.String())
	return err
}
