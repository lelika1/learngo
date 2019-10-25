package offset_test

import (
	"fmt"
	"testing"

	"github.com/lelika1/learngo/internal/offset"
)

func TestToLineColumn(t *testing.T) {
	const input = "abcd\nef\ngt"
	m := offset.NewMapper(input)

	tests := []struct {
		offset int
		// Format: "line:col" for successful result; "?" for failed one.
		want string
	}{
		{0, "0:0"},
		{1, "0:1"},
		{3, "0:3"},
		{4, "0:4"},
		{5, "1:0"},
		{6, "1:1"},
		{9, "2:0"},
		{99, "?"},
		{-1, "?"},
	}
	for _, test := range tests {
		line, col, ok := m.ToLineColumn(test.offset)
		if !ok {
			if test.want != "?" {
				t.Errorf("ToLineColumn(%v) = (%v:%v, %v), want %v",
					test.offset, line, col, ok, test.want)
			}
			continue
		}
		if got := fmt.Sprintf("%v:%v", line, col); got != test.want {
			t.Errorf("ToLineColumn(%v) = (%v:%v, %v), want %v", test.offset, line, col, ok, test.want)
		}
	}
}

func TestLineOffset(t *testing.T) {
	const input = "abcd\nef\ngt"
	m := offset.NewMapper(input)

	tests := []struct {
		line int
		want int // -1 the call should fail
	}{
		{line: 0, want: 0},
		{line: 1, want: 5},
		{line: 2, want: 9},
		{line: -1, want: -1},
		{line: 3, want: -1},
		{line: 4, want: -1},
	}
	for _, test := range tests {
		got, ok := m.LineOffset(test.line)
		if (!ok && test.want != -1) || (ok && got != test.want) {
			t.Errorf("LineOffset(%v) = (%v, %v), want %v", test.line, got, ok, test.want)
		}
	}
}
