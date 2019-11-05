package canvas_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lelika1/learngo/internal/canvas"
)

func pict(row string, rows ...string) string {
	var sb strings.Builder
	sb.WriteString(row)
	sb.WriteRune('\n')

	cols := len(row)
	for i, r := range rows {
		if len(r) != cols {
			panic(fmt.Sprintf("Line %v len = %v, want %v", i, len(r), cols))
		}
		sb.WriteString(r)
		sb.WriteRune('\n')
	}
	return sb.String()
}

func TestNewCanvas(t *testing.T) {
	tests := []struct {
		w int
		h int
		// "?" in case of nil
		want string
	}{
		{2, 3, "..\n..\n..\n"},
		{1, 3, ".\n.\n.\n"},
		{3, 1, "...\n"},
		{4, 0, "?"},  //wrong size
		{-2, 3, "?"}, //wrong size
	}

	for _, tc := range tests {
		c, ok := canvas.NewCanvas(tc.w, tc.h)
		if !ok {
			if tc.want != "?" {
				t.Errorf("NewCanvas(%v, %v) = %v, want %q", tc.w, tc.h, nil, tc.want)
			}
			continue
		}

		got := c.Export(canvas.Color.String)
		if got != tc.want {
			t.Errorf("NewCanvas(%v, %v) = %q, want %q", tc.w, tc.h, got, tc.want)
		}
	}
}

func TestImportExport(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   string
	}{
		{
			pict(
				"..y..",
				".yry.",
				"yyxyy",
				".yyb.",
				"..g..",
			),
			pict(
				"..y..",
				".yry.",
				"yyxyy",
				".yyb.",
				"..g..",
			),
			"",
		},

		// different errors
		{
			"",
			"",
			"canvas should have positive dimensions",
		},
		{
			"\n",
			"",
			"canvas should have positive dimensions",
		},
		{
			"\n\n",
			"",
			"canvas should have positive dimensions",
		},
		{
			"x.\n..\n..",
			"",
			"should be \\n in the end of file",
		},
		{

			"..y..\n.yy.\nyyyyy\n.yyy.\n..y..\n",
			"",
			"wrong length of row 1 in canvas",
		},
		{
			pict(
				"..y..",
				".yyh.",
				"yyyyy",
				".yyy.",
				"..y..",
			),
			"",
			"wrong color at [3:1] = h",
		},
	}

	for _, tc := range tests {
		c, err := canvas.Import(tc.input)

		if (err != nil && tc.err != err.Error()) || (tc.err != "" && err == nil) {
			t.Errorf("Import(%q) return status %v, want %v", tc.input, err.Error(), tc.err)
		}
		if err == nil {
			if got := c.Export(canvas.Color.String); got != tc.want {
				t.Errorf("Import/Export(%q) = %q, want %q", tc.input, got, tc.want)
			}
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		input canvas.Point
		color canvas.Color
		want  string
	}{
		{
			canvas.Point{0, 0}, canvas.Black,
			pict(
				"x.",
				"..",
				"..",
			),
		},
		{
			canvas.Point{1, 2}, canvas.Green,
			pict(
				"..",
				"..",
				".g",
			),
		},
		{
			canvas.Point{1, 2}, canvas.Blue,
			pict(
				"..",
				"..",
				".b",
			),
		},
		{ // out of bounds
			canvas.Point{-1, 2}, canvas.Green,
			pict(
				"..",
				"..",
				"..",
			),
		},
		{ // out of bounds
			canvas.Point{2, 3}, canvas.Green,
			pict(
				"..",
				"..",
				"..",
			),
		},
		{ // out of bounds
			canvas.Point{2, 1}, canvas.Black,
			pict(
				"..",
				"..",
				"..",
			),
		},
	}

	for _, tc := range tests {
		c, _ := canvas.NewCanvas(2, 3)
		c.Dot(tc.input, tc.color)
		got := c.Export(canvas.Color.String)
		if got != tc.want {
			t.Errorf("Dot(%v, %v) = %q, want %q", tc.input, tc.color, got, tc.want)
		}
	}
}

func TestRect(t *testing.T) {
	tests := []struct {
		left  canvas.Point
		right canvas.Point
		color canvas.Color
		want  string
	}{
		{
			canvas.Point{1, 0}, canvas.Point{2, 3}, canvas.Red,
			pict(
				".rr",
				".rr",
				".rr",
				".rr",
			),
		},
		{ // rect on all canvas
			canvas.Point{0, 0}, canvas.Point{2, 3}, canvas.Black,
			pict(
				"xxx",
				"xxx",
				"xxx",
				"xxx",
			),
		},
		{ // bigger than canvas
			canvas.Point{-1, -1}, canvas.Point{2, 3}, canvas.Black,
			pict(
				"xxx",
				"xxx",
				"xxx",
				"xxx",
			),
		},
		{ // rect in shape of line
			canvas.Point{0, 0}, canvas.Point{0, 2}, canvas.Blue,
			pict(
				"b..",
				"b..",
				"b..",
				"...",
			),
		},
		{ // out of bounds
			canvas.Point{-2, -3}, canvas.Point{-1, -1}, canvas.Black,
			pict(
				"...",
				"...",
				"...",
				"...",
			),
		},
		{ // left and right corners are swaped
			canvas.Point{2, 3}, canvas.Point{0, 0}, canvas.Black,
			pict(
				"...",
				"...",
				"...",
				"...",
			),
		},
	}

	for _, tc := range tests {
		c, _ := canvas.NewCanvas(3, 4)
		c.Rect(tc.left, tc.right, tc.color)
		got := c.Export(canvas.Color.String)
		if got != tc.want {
			t.Errorf("Rect(%v, %v, %v) = %q, want %q", tc.left, tc.right, tc.color, got, tc.want)
		}
	}
}

func TestCircle(t *testing.T) {
	tests := []struct {
		center canvas.Point
		rad    int
		color  canvas.Color
		want   string
	}{
		{
			canvas.Point{2, 2}, 2, canvas.Yellow,
			pict(
				"..y..",
				".yyy.",
				"yyyyy",
				".yyy.",
				"..y..",
			),
		},
		{
			canvas.Point{3, 3}, 1, canvas.Blue,
			pict(
				".....",
				".....",
				"...b.",
				"..bbb",
				"...b.",
			),
		},
		{ // shifted circle
			canvas.Point{0, 2}, 2, canvas.Red,
			pict(
				"r....",
				"rr...",
				"rrr..",
				"rr...",
				"r....",
			),
		},
		{ // very big radius
			canvas.Point{2, 2}, 10, canvas.Green,
			pict(
				"ggggg",
				"ggggg",
				"ggggg",
				"ggggg",
				"ggggg",
			),
		},
		{ // zero radius
			canvas.Point{2, 2}, 0, canvas.Green,
			pict(
				".....",
				".....",
				".....",
				".....",
				".....",
			),
		},
		{ // out of bounds
			canvas.Point{-3, -3}, 2, canvas.Red,
			pict(
				".....",
				".....",
				".....",
				".....",
				".....",
			),
		},
	}

	for _, tc := range tests {
		c, _ := canvas.NewCanvas(5, 5)
		c.Circle(tc.center, tc.rad, tc.color)
		got := c.Export(canvas.Color.String)
		if got != tc.want {
			t.Errorf("Circle(%v, %v, %v) = %q, want %q", tc.center, tc.rad, tc.color, got, tc.want)
		}
	}
}

func TestFill(t *testing.T) {
	input := pict(
		"gggb.",
		"gggb.",
		"gggb.",
		"bbbr.",
		".....",
	)

	tests := []struct {
		p     canvas.Point
		color canvas.Color
		want  string
	}{
		{
			canvas.Point{3, 3}, canvas.Black,
			pict(
				"gggb.",
				"gggb.",
				"gggb.",
				"bbbx.",
				".....",
			),
		},
		{
			canvas.Point{3, 3}, canvas.White,
			pict(
				"gggb.",
				"gggb.",
				"gggb.",
				"bbb..",
				".....",
			),
		},
		{
			canvas.Point{1, 1}, canvas.Black,
			pict(
				"xxxb.",
				"xxxb.",
				"xxxb.",
				"bbbr.",
				".....",
			),
		},
		{
			canvas.Point{3, 2}, canvas.Black,
			pict(
				"gggx.",
				"gggx.",
				"gggx.",
				"bbbr.",
				".....",
			),
		},
		{
			canvas.Point{4, 1}, canvas.Black,
			pict(
				"gggbx",
				"gggbx",
				"gggbx",
				"bbbrx",
				"xxxxx",
			),
		},
		{ // the same color to fill
			canvas.Point{0, 0}, canvas.Green,
			pict(
				"gggb.",
				"gggb.",
				"gggb.",
				"bbbr.",
				".....",
			),
		},
		{ // out of bounds
			canvas.Point{-1, 0}, canvas.Black,
			input,
		},
	}

	for _, tc := range tests {
		c, err := canvas.Import(input)
		if err != nil {
			t.Fatalf("canvas.Import(%q): %v", input, err)
		}
		c.Fill(tc.p, tc.color)
		got := c.Export(canvas.Color.String)
		if got != tc.want {
			t.Errorf("Fill(%v, %v) = %q, want %q", tc.p, tc.color, got, tc.want)
		}
	}
}
