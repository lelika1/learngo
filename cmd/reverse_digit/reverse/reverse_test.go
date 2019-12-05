package reverse_test

import (
	"testing"

	"github.com/lelika1/learngo/cmd/reverse_digit/reverse"
)

func TestReverseDigit(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{1234, 4321},
		{12, 21},
		{10, 1},
		{55, 55},
		{1, 1},
		{0, 0},
		{-1, -1},
		{-90, -9},
		{-253, -352},
	}

	for _, ts := range tests {
		got := reverse.Digit(ts.number)
		if got != ts.want {
			t.Errorf("Digit(%v) = %v, want %v", ts.number, got, ts.want)
		}
	}
}
