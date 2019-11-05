package canvas

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Canvas represents a 2D surface of a particular size.
// Coordinates on this surface are zero-based.
// Point (x:0, y:0) is in the top left corner.
// Horizontal axis corresponds to X coordinate (grows to the right).
// Vertical axis corresponds to Y coordinate (grows to the bottom).
//
// Example of a 2x3 canvas:
// |---------------|
// | (0,0) | (1,0) |
// |---------------|
// | (0,1) | (1,1) |
// |---------------|
// | (0,2) | (1,2) |
// |---------------|
type Canvas struct {
	canvas [][]Color
	hight  int
	width  int
}

// Import creates Canvas given an output of some canvas.Export() function.
func Import(s string) (*Canvas, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("canvas should have positive dimensions")
	}

	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	} else {
		return nil, fmt.Errorf("should be \\n in the end of file")
	}

	rows := strings.Split(s, "\n")
	if len(rows) == 0 || len(rows[0]) == 0 {
		return nil, fmt.Errorf("canvas should have positive dimensions")
	}

	width := len(rows[0])
	for i, row := range rows {
		if len(row) != width {
			return nil, fmt.Errorf("wrong length of row %v in canvas", i)
		}
	}

	c, _ := NewCanvas(width, len(rows))
	for y, row := range rows {
		for x, sym := range row {
			if cl, ok := parseColor(sym); ok {
				c.Dot(Point{X: x, Y: y}, cl)
			} else {
				return nil, fmt.Errorf("wrong color at [%v:%v] = %c", x, y, sym)
			}
		}
	}
	return c, nil
}

// NewCanvas allocates a new canvas with the given hight and width.
// hight and width must be bigger than zero
func NewCanvas(w, h int) (*Canvas, bool) {
	if w <= 0 || h <= 0 {
		return nil, false
	}

	c := Canvas{hight: h, width: w}

	c.canvas = make([][]Color, h)
	pixels := make([]Color, h*w)
	for i := range c.canvas {
		c.canvas[i], pixels = pixels[:w], pixels[w:]
	}
	return &c, true
}

func (c *Canvas) inBounds(p Point) bool {
	return p.X >= 0 && p.X < c.width && p.Y >= 0 && p.Y < c.hight
}

// get returns color of given point
func (c *Canvas) get(p Point) Color {
	return c.canvas[p.Y][p.X]
}

// Dot paints the given point with the given color.
func (c *Canvas) Dot(p Point, cl Color) {
	if c.inBounds(p) {
		c.canvas[p.Y][p.X] = cl
	}
}

// Circle draws a circle with the given center and radius.
func (c *Canvas) Circle(center Point, r int, cl Color) {
	if r <= 0 {
		return
	}

	for y, row := range c.canvas {
		for x := range row {
			if (center.X-x)*(center.X-x)+(center.Y-y)*(center.Y-y) <= r*r {
				c.Dot(Point{X: x, Y: y}, cl)
			}
		}
	}
}

// Rect draws a rectangle given its top-left and bottom-right corners.
func (c *Canvas) Rect(left, right Point, cl Color) {
	if left.X > right.X || left.Y > right.Y {
		return
	}

	for y := left.Y; y <= right.Y; y++ {
		for x := left.X; x <= right.X; x++ {
			c.Dot(Point{X: x, Y: y}, cl)
		}
	}
}

// Fill fills the area with the given Point and color.
func (c *Canvas) Fill(p Point, cl Color) {
	if !c.inBounds(p) {
		return
	}

	origCol := c.get(p)
	if origCol == cl {
		return // the same color - nothing to color
	}

	var area []Point
	addPoint := func(p Point) {
		if c.inBounds(p) && c.get(p) == origCol {
			c.Dot(p, cl)
			area = append(area, p)
		}
	}

	addPoint(p) //first point
	for len(area) != 0 {
		p = area[len(area)-1]
		area = area[:len(area)-1]

		addPoint(Point{p.X + 1, p.Y}) // right point
		addPoint(Point{p.X - 1, p.Y}) // left point
		addPoint(Point{p.X, p.Y - 1}) // up point
		addPoint(Point{p.X, p.Y + 1}) // down point
	}
}

// Format ...
type Format func(c Color) string

// Export returns a string represantation of the canvas with given format
func (c *Canvas) Export(f Format) string {
	var sb strings.Builder

	for _, row := range c.canvas {
		for _, p := range row {
			sb.WriteString(f(p))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

// Point describes coordinates of one pixel in the Canvas.
type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return "{" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + "}"
}

// Color of a pixel in the Canvas.
type Color int8

// All possible colors.
const (
	White Color = iota
	Black
	Red
	Green
	Blue
	Yellow

	MaxColor
)

// String ...
func (c Color) String() string {
	switch c {
	case White:
		return "."
	case Black:
		return "x"
	case Red:
		return "r"
	case Green:
		return "g"
	case Blue:
		return "b"
	case Yellow:
		return "y"
	}
	return "?"
}

// ColoredString ...
func ColoredString(c Color) string {
	switch c {
	case White:
		return color.WhiteString("█")
	case Black:
		return color.BlackString("█")
	case Red:
		return color.RedString("█")
	case Green:
		return color.GreenString("█")
	case Blue:
		return color.BlueString("█")
	case Yellow:
		return color.YellowString("█")
	}
	return "?"
}

func parseColor(s rune) (Color, bool) {
	switch s {
	case '.':
		return White, true
	case 'x':
		return Black, true
	case 'r':
		return Red, true
	case 'g':
		return Green, true
	case 'b':
		return Blue, true
	case 'y':
		return Yellow, true
	}
	return White, false
}
