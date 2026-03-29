//revive:disable:package-comments
package structured

// Options configures the structured handler's hierarchical attribute rendering.
// Level and color are always read from environment (LOG_LEVEL, NO_COLOR).
type Options struct {
	// AttributeLevels defines the hierarchical grouping of log attributes.
	// Each nested array represents a nesting level:
	//   - Level 0 (empty): message line only, no inline attributes
	//   - Level 1 (2 spaces): service-level context and source location
	//   - Level 2 (4 spaces): span context
	//   - Level 3 (6 spaces): operation details
	//   - Level 4+ (8+ spaces): unknown attributes, sorted alphabetically
	// If nil, uses DefaultAttributeLevels from logger package.
	AttributeLevels [][]string
}

// DefaultAttributeLevels defines the default hierarchical structure for structured logging.
// Each nested array represents a nesting level:
//   - Level 0 (empty): message line only, no inline attributes
//   - Level 1 (2 spaces): service-level context and source location
//   - Level 2 (4 spaces): span context
//   - Level 3 (6 spaces): operation details
//   - Level 4+ (8+ spaces): unknown attributes, sorted alphabetically
var DefaultAttributeLevels = [][]string{
	{},
	{"source", "service", "address", "component", "trace_id"},
	{"span_id"},
	{"method"},
}
