/* This code snippet was taken from the Go Tour.
 *
 * It shows how easy it is to do simple multithreaded tasks in Go.
 *
 * This particular example fires up a pair of goroutines that concurrently
 * sum each half of a slice and write the result into a channel. The results
 * are read from the channel in the main thread and combined together to produce
 * the final result.
 *
 * It is educational and very interesting to see how simple and brief the code is when
 * compared to, say, C or C++
 *
 */

package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 6, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <- c, <- c
	fmt.Println("Sum =", x+y)
}
