/* Taken from the Go Tour (Exercise: Errors)
 *
 * Copy your Sqrt function from the earlier exercise and modify it to return an error value.
 *
 * Sqrt should return a non-nil error value when given a negative number, as it doesn't support
 * complex numbers.
 *
 * Create a new type
 *
 * type ErrNegativeSqrt float64
 *
 * and make it an error by giving it a
 *
 * func (e ErrNegativeSqrt) Error() string
 *
 * method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
 *
 * Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite
 * loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
 *
 * Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
 *
 */

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

/* We need to convert `e' back to a float64 here to avoid entering an infinite loop.
 * If we don't, the code ends up calling ErrNegativeSqrt.Error() to convert `e' into a string, which
 * is what the function is supposed to do in the first place!
 */
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
