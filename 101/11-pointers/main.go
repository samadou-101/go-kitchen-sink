package main

import "fmt"

func zeroVal(v int) {
	v = 0
}

func zeroPtr(v *int) {
	*v = 0
}

func main() {
	x := 10

	// & gets address, * dereferences
	p := &x
	fmt.Println(*p)

	*p = 20
	fmt.Println(x)

	// pass by value vs pointer
	zeroVal(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)
}
