package main

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/lelika1/learngo/internal/canvas"
)

// NewPalette creates new color.Palette coresponding to canvas.Color.
func NewPalette() *color.Palette {
	var palette color.Palette
	palette = make([]color.Color, 8)
	palette[canvas.White] = color.RGBA{255, 255, 255, 1}
	palette[canvas.Black] = color.RGBA{0, 0, 0, 1}
	palette[canvas.Red] = color.RGBA{255, 0, 0, 1}
	palette[canvas.Green] = color.RGBA{0, 255, 0, 1}
	palette[canvas.Blue] = color.RGBA{0, 0, 50, 1}
	palette[canvas.Yellow] = color.RGBA{255, 255, 0, 1}
	palette[canvas.Magenta] = color.RGBA{255, 0, 255, 1}
	palette[canvas.Cyan] = color.RGBA{0, 255, 255, 1}
	return &palette
}

// NearestColor finds nearest ccolor from canves.Color.
func NearestColor(p *color.Palette, pix color.Color) canvas.Color {
	col := p.Index(pix)
	return canvas.Color(col)
}

func main() {
	imgfile, err := os.Open("./cmd/jpg_reader/input.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgfile.Close()

	img, err := jpeg.Decode(imgfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	c, ok := canvas.NewCanvas(bounds.Max.X, bounds.Max.Y)
	if !ok {
		fmt.Println("NewCanvas failed")
		os.Exit(1)
	}

	p := NewPalette()
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			col := NearestColor(p, img.At(i, j))
			c.Dot(canvas.Point{X: i, Y: j}, col)
		}
	}
	fmt.Println(c.Export(canvas.ColoredString))
}
