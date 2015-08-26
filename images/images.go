package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	colorAt func(x, y int) uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	c := i.colorAt(x, y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	/* Choose the color function by changing the closure below
         * Interesting functions include x^y, x*y and (x+y)/2
         */
	m := Image{257, 257, func (x, y int) uint8 { return uint8(x*y) }}
	pic.ShowImage(m)
}
