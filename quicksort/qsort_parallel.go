
package main

import "fmt"

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func quicksort_parallel(arr []int, c chan int) {

	if (len(arr) <= 1) {
		c <- 1
		return
	}

	swap(arr, 0, len(arr)/2)

	i := 0
	for j := 1; j < len(arr); j++ {
		if (arr[j] <= arr[0]) {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, 0, i)

	c1, c2 := make(chan int, 1), make(chan int, 1)

	go quicksort_parallel(arr[:i], c1)
	go quicksort_parallel(arr[i+1:], c2)

	_ = <- c1
	_ = <- c2
	c <- 1
}

func Quicksort(arr []int) {
	c := make(chan int, 1)
	go quicksort_parallel(arr, c)
	_ = <- c
}

func main() {
	tests := [][]int{
		{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 },
		{ 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 },
		{ 3, 0, 9, 9, 0, -3, -1, 0, -3, 10, 99, -3 -5 },
		{ 1 },
		{ 1, 2 },
		{ 2, 1 },
		{ 1, 2, 3 },
		{ 1, 3, 2 },
		{ 2, 1, 3 },
		{ 2, 3, 1 },
		{ 3, 1, 2 },
		{ 3, 2, 1},
	}

	for _, a := range tests {
		Quicksort(a)
		fmt.Println(a)
	}

}
