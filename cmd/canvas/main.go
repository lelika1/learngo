package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/canvas"
)

func main() {
	c, ok := canvas.NewCanvas(50, 50)
	if !ok {
		return
	}

	for i := 0; i < 25; i++ {
		cl := canvas.Color(i % int(canvas.MaxColor))
		c.Rect(canvas.Point{X: i, Y: i}, canvas.Point{X: 49 - i, Y: 49 - i}, cl)
	}
	for i := 24; i > 0; i-- {
		cl := canvas.Color(i % int(canvas.MaxColor))
		c.Circle(canvas.Point{X: 25, Y: 25}, i, cl)
	}

	fmt.Println(c.ExportColor())
}
