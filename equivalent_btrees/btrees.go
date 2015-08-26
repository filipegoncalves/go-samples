package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func walkAux(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkAux(t.Left, ch)
	ch <- t.Value
	walkAux(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	walkAux(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch_t1 := make(chan int)
	ch_t2 := make(chan int)

	go Walk(t1, ch_t1)
	go Walk(t2, ch_t2)

	t1_v, t1_ok := <- ch_t1
	t2_v, t2_ok := <- ch_t2

	for t1_ok && t2_ok && t1_v == t2_v {
		t1_v, t1_ok = <- ch_t1
		t2_v, t2_ok = <- ch_t2
	}

	return !t1_ok && !t2_ok
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Same(tree.New(1), tree.New(1)) =", Same(tree.New(1), tree.New(1)))
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Same(tree.New(1), tree.New(2)) =", Same(tree.New(1), tree.New(2)))
	}
}
