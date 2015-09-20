
package main

import (
	"fmt"
)

func say(right, left chan int) {
	left <- 1 + <- right
}

func main() {
	const gophers = 100000
	rightmost := make(chan int)
	right := rightmost
	var left chan int

	for i := 0; i < gophers; i++ {
		left = make(chan int)
		go say(right, left)
		right = left
	}

	go func() { rightmost <- 1 }()

	fmt.Println(<- left)
}
