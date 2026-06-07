package main

import "fmt"

// constants = compile-time fixed values
// no :=
// can be typed or untyped
// iota = enum generator
// more flexible than variables in expressions

func main() {
	const num = 10 // no type assigned until used
	const num2 int = 20

	const (
		a = iota
		b
		c
	)

	fmt.Println(a, b, c)
}
