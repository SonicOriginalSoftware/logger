//revive:disable:package-comments
package attrs

// Level stores the hierarchical level information for an attribute
type Level struct {
	Level int // Which indent level (0=line1, 1=2spaces, 2=4spaces, etc.)
	Order int // Order within the level for consistent rendering
}

// BuildLevelMap creates a lookup map from Levels configuration
// This is called once during handler creation for O(1) attribute lookups
func BuildLevelMap(levels [][]string) map[string]Level {
	// Count total attributes for map pre-allocation
	count := 0
	for _, attrs := range levels {
		count += len(attrs)
	}

	// Pre-allocate map with known capacity for better performance
	attrMap := make(map[string]Level, count)
	for levelIdx, attrs := range levels {
		for orderIdx, attrName := range attrs {
			attrMap[attrName] = Level{
				Level: levelIdx,
				Order: orderIdx,
			}
		}
	}
	return attrMap
}
