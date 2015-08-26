/* Taken from the Go Tour (Exercise: Images)
 *
 * Remember the picture generator you wrote earlier? Let's write another one, but this time it will
 * return an implementation of image.Image instead of a slice of data.
 *
 * Define your own Image type, implement the necessary methods, and call pic.ShowImage.
 *
 * Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
 *
 * ColorModel should return color.RGBAModel.
 *
 * At should return a color; the value v in the last picture generator corresponds to
 * color.RGBA{v, v, 255, 255} in this one.
 */

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
