/* Taken from the Go Tour (Exercise: Slices)
 *
 * Implement Pic. It should return a slice of length dy, each element of which is a slice
 * of dx 8-bit unsigned integers. When you run the program, it will display your picture,
 * interpreting the integers as grayscale (well, bluescale) values.
 *
 * The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
 *
 * (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
 *
 * (Use uint8(intValue) to convert between types.)
 *
 */

package main

import "golang.org/x/tour/pic"

/*
 * Choose here the function to generate the picture
 * You can choose between:
 *
 * pic1: (x+y)/2
 * pic2: x*y
 * pic3: x^y
 *
 */

var pic_fun = pic1

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
		for j := range picture[i] {
			picture[i][j] = pic_fun(j, i)
		}
	}
	return picture
}

func pic1(x, y int) uint8 {
	return (uint8(x) + uint8(y)) / 2
}

func pic2(x, y int) uint8 {
	return uint8(x) * uint8(y)
}

func pic3(x, y int) uint8 {
	return uint8(x) ^ uint8(y)
}
