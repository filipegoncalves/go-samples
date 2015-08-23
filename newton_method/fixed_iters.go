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
