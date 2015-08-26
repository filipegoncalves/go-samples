/* Taken from the Go Tour (Exercise: Equivalent Binary Trees)
 *
 * There can be many different binary trees with the same sequence of values stored at the leaves.
 * For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.
 *
 *                                    8
 *         3                         / \
 *       /   \                      3   13
 *      /     \                   /   \
 *     /       \                 /     \
 *    1         8               1       5
 *  /   \     /   \           /   \
 * 1     2   5    13         1     2
 *
 * A function to check whether two binary trees store the same sequence is quite complex in most
 * languages. We'll use Go's concurrency and channels to write a simple solution.
 *
 * This example uses the tree package, which defines the type:
 *
 * type Tree struct {
 *    Left  *Tree
 *    Value int
 *    Right *Tree
 * }
 *
 * 1. Implement the Walk function.
 *
 * 2. Test the Walk function.
 *
 * The function tree.New(k) constructs a randomly-structured binary tree holding the
 * values k, 2k, 3k, ..., 10k.
 *
 * Create a new channel ch and kick off the walker:
 *
 * go Walk(tree.New(1), ch)
 *
 * Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
 *
 * 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
 *
 * 4. Test the Same function.
 *
 * Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should
 * return false.
 *
 */

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
