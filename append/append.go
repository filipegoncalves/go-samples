/* A small example demonstrating how append() is implemented
 *
 * Based on the ideas exposed in http://blog.golang.org/slices
 */

package main

import "fmt"

var allocs int

func my_append(slice []int, items ...int) []int {
	total := len(slice)+len(items)

	if total > cap(slice) {
		newSlice := make([]int, len(slice), total*2+1)
		copy(newSlice, slice)
		slice = newSlice
		allocs++
	}

	n := len(slice)
	slice = slice[:total]
	copy(slice[n:], items)
	return slice
}

func main() {
	var s []int

	for i := 0; i < 20; i++ {
		fmt.Println("i =", i, "s =", s)
		s = my_append(s, i)
	}

	fmt.Println("Finished, s =", s)
	s = my_append(s, s...)
	fmt.Println("After appending to itself:", s)
	fmt.Println("Total allocations:", allocs)
}
