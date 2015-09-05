
package main

import "fmt"

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Quicksort(arr []int) {

	if (len(arr) <= 1) {
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

	Quicksort(arr[:i])
	Quicksort(arr[i+1:])
}

func main() {
	//test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//test := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	test := []int{3, 0, 9, 9, 0, -3, -1, 0, -3, 10, 99, -3 -5}
	Quicksort(test)
	fmt.Println(test)
}
