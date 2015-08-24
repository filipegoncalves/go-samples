
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	if x == 0 {
		return 0, nil
	}

	z := 1.0
	prev := 0.0
	const delta = 0.00000000001

	for math.Abs(z-prev) > delta {
		prev = z
		z -= (z*z-x)/(2*z)
	}

	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	for i := -5; i < 10; i++ {
		res, err := Sqrt(float64(i))
		if err != nil {
			fmt.Println("Sqrt error:", err)
		} else {
			fmt.Printf("Sqrt(%v) = %v\n", i, res)
		}
	}
}
