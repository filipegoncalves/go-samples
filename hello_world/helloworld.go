/* A clever way of printing Hello World using Go's "defer" capability
 * It's nothing extraordinary, but it's funny when you learn it for the first time!
 */

package main

import "fmt"

func main() {
	defer fmt.Println("!")
	defer fmt.Printf("World")
	fmt.Printf("Hello, ")
}
