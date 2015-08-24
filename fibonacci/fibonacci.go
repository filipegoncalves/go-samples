/* Taken from the Go Tour (Exercise: Fibonacci closure)
 *
 * Let's have some fun with functions.
 *
 * Implement a fibonacci function that returns a function (a closure) that returns successive
 * fibonacci numbers.
 *
 */

package main

import "fmt"

func fibonacci() func() int {
	before_last, last := -1, 1
	return func() int {
		last, before_last = before_last+last, last
		return last
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
