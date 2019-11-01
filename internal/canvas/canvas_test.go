package canvas_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/canvas"
)

func TestDot(t *testing.T) {
	tests := []struct {
		input canvas.Point
		color canvas.Color
		want  string
	}{
		{
			canvas.Point{0, 0}, canvas.Black,
			"x \n  \n  ",
		},
		{
			canvas.Point{1, 2}, canvas.Green,
			"  \n  \n g",
		},
	}

	for _, tc := range tests {
		c := canvas.NewCanvas(2, 3)
		c.Dot(tc.input, tc.color)
		got := c.Export()
		if got != tc.want {
			t.Errorf("Dot(%v:%v, %v) = %#v, want %#v", tc.input.X, tc.input.Y, tc.color, got, tc.want)
		}
	}
}
