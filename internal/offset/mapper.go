// Package offset contains utilities for mapping byte offsets to line:col in content.
package offset

// Mapper is built for a particular content and provides
// fast conversion between byte offsets and line:column.
//
// Our assumptions:
// * Lines are separated by '\n' symbols.
// * Line and column are zero-based.
// * There are no unicode characters in the given text.
type Mapper struct {
	// TODO: Implement.
}

// NewMapper returns a mapper for a given content.
func NewMapper(content string) *Mapper {
	// TODO: Implement.
	return nil
}

// ToLineColumn returns zero-based line and column for a given byte offset.
// For offset that is out-of-bound returns ok=false.
func (m *Mapper) ToLineColumn(offset int) (line, col int, ok bool) {
	// TODO: Implement.
	return
}

// LineOffset returns byte offset for the begining of the given string.
func (m *Mapper) LineOffset(line int) (offset int, ok bool) {
	// TODO: Implement.
	return
}
