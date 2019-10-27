package offset_test

import (
	"fmt"
	"testing"

	"github.com/lelika1/learngo/internal/offset"
)

func TestToLineColumn(t *testing.T) {
	const input = "abcd\nef\ngt\n"
	tests := []struct {
		input  string
		offset int
		// Format: "line:col" for successful result; "?" for failed one.
		want string
	}{
		{input, 0, "0:0"},
		{input, 1, "0:1"},
		{input, 3, "0:3"},
		{input, 4, "0:4"},
		{input, 5, "1:0"},
		{input, 6, "1:1"},
		{input, 8, "2:0"},
		{input, 9, "2:1"},
		{input, 10, "2:2"},
		{input, 11, "?"},
		{input, 99, "?"},
		{input, -1, "?"},
		// Testing empty input
		{input: "", offset: 0, want: "?"},
		{input: "", offset: 5, want: "?"},
	}
	for _, test := range tests {
		m := offset.NewMapper(test.input)
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
	const input = "abcd\nef\ngt\n\na"
	tests := []struct {
		input string
		line  int
		// Format: "offset:size" for successful result; "?" for failed one.
		want string
	}{
		{input, 0, "0:5"},
		{input, 1, "5:3"},
		{input, 2, "8:3"},
		{input, -1, "?"},
		{input, 3, "11:1"},
		{input, 4, "12:1"},
		{input, 5, "?"},
		// Testing empty input
		{input: "", line: 0, want: "?"},
		{input: "", line: 5, want: "?"},
	}
	for _, test := range tests {
		m := offset.NewMapper(test.input)
		lineOffset, lineSize, ok := m.LineOffset(test.line)
		if !ok {
			if test.want != "?" {
				t.Errorf("LineOffset(%v) = (%v:%v, %v), want %v",
					test.line, lineOffset, lineSize, ok, test.want)
			}
			continue
		}

		if got := fmt.Sprintf("%v:%v", lineOffset, lineSize); got != test.want {
			t.Errorf("LineOffset(%v) = (%v, %v), want %v", test.line, got, ok, test.want)
		}
	}
}
