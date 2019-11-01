package canvas

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
	// TODO: implement.
}

// NewCanvas allocates a new canvas with the given hight and width.
func NewCanvas(h, w int) *Canvas {
	// TODO: implement.
	return nil
}

// Dot paints the given point with the given color.
func (c *Canvas) Dot(p Point, cl Color) {
	// TODO: implement.
}

// Circe draws a circle with the given center and radius.
// FIXME: I'm not sure if radius should be float64. Maybe int will be enough?
func (c *Canvas) Circe(center Point, r float64, cl Color) {
	// TODO: implement.
}

// Rect draws a rectangle given its top-left and bottom-right corners.
func (c *Canvas) Rect(a, b Point, cl Color) {
	// TODO: implement.
}

// Export retuns a string representation of the Canvas.
// TODO: Advanced task - have a way to print colored output to terminal.
func (c *Canvas) Export() string {
	// TODO: implement.
	return ""
}

// Point describes coordinates of one pixel in the Canvas.
type Point struct {
	X int
	Y int
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
)

func (c Color) String() string {
	switch c {
	case White:
		return " "
	case Black:
		return "x"
	case Red:
		return "r"
	case Green:
		return "g"
	case Blue:
		return "b"
	}
	return "?"
}
