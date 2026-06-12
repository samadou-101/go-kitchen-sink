package main

import "fmt"

func main() {
	// fixed-size arrays
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	// array literal
	b := [3]string{"go", "rust", "zig"}
	fmt.Println(b)

	// ellipsis counts elements
	c := [...]bool{true, false, true}
	fmt.Println(c)

	// len and index
	fmt.Println(len(c), c[1])
}
