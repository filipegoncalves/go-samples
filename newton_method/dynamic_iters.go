package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {

	if x == 0 {
		return 0
	}

	z = 1.0
	prev := 0.0
	const delta = 0.00000000001

	for math.Abs(z-prev) > delta {
		prev = z
		z -= (z*z-x)/(2*z)
	}

	return
}

func print_sqrt(n int) {
	fmt.Printf("sqrt(%v) = %v\n", n, Sqrt(float64(n)))
	fmt.Printf("math.Sqrt(%v) = %v\n", n, math.Sqrt(float64(n)))
}

func main() {
	for i := 0; i < 10; i++ {
		print_sqrt(i)
	}
}
