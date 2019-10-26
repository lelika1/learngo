// Package offset contains utilities for mapping byte offsets to line:col in content.
package offset

import (
	"sort"
	"strings"
)

// Mapper is built for a particular content and provides
// fast conversion between byte offsets and line:column.
//
// Our assumptions:
// * Lines are separated by '\n' symbols.
// * Line and column are zero-based.
// * There are no unicode characters in the given text.
type Mapper struct {
	rowStart    []int
	contentSize int
}

// NewMapper returns a mapper for a given content.
func NewMapper(content string) *Mapper {
	rows := strings.SplitAfter(content, "\n")
	rowStart := make([]int, 1, len(rows))
	for i := 1; i < len(rows); i++ {
		nextRowStart := rowStart[i-1] + len(rows[i-1])
		if nextRowStart < len(content) {
			rowStart = append(rowStart, nextRowStart)
		}
	}

	return &Mapper{
		rowStart:    rowStart,
		contentSize: len(content),
	}
}

// ToLineColumn returns zero-based line and column for a given byte offset.
// For offset that is out-of-bound returns ok=false.
func (m *Mapper) ToLineColumn(offset int) (line, col int, ok bool) {
	if offset < 0 || offset >= m.contentSize {
		return 0, 0, false
	}

	row := sort.SearchInts(m.rowStart, offset)
	if row == len(m.rowStart) || m.rowStart[row] > offset {
		row--
	}
	return row, offset - m.rowStart[row], true
}

// LineOffset returns byte offset for the begining of the given string.
func (m *Mapper) LineOffset(line int) (offset int, ok bool) {
	if line < 0 || line >= len(m.rowStart) {
		return 0, false
	}
	return m.rowStart[line], true
}
