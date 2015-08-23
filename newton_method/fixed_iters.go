/* Taken from the Go Tour (Exercise: Loops and Functions)
 *
 * As a simple way to play with functions and loops, implement the square root function using
 * Newton's method.
 *
 * In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then
 * repeating:
 *
 * z = z - (z*z-x)/(2*z)
 *
 * To begin with, just repeat that calculation 10 times and see how close you get to the answer for
 * various values (1, 2, 3, ...).
 *
 * Next, change the loop condition to stop once the value has stopped changing (or only changes by
 * a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?
 *
 * Hint: to declare and initialize a floating point value, give it floating point syntax or use a
 * conversion:
 *
 * z := float64(1)
 * z := 1.0
 *
 */

package main

import "fmt"

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z-x)/(2*z)
	}
	return
}

func print_sqrt(n int) {
	fmt.Printf("sqrt(%v) = %v\n", n, Sqrt(float64(n)))
}

func main() {
	for i := 0; i < 10; i++ {
		print_sqrt(i)
	}
}
