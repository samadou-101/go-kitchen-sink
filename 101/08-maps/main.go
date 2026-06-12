package main

import "fmt"

func main() {
	// make
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// map literal
	m2 := map[string]string{
		"go":   "Golang",
		"rust": "Rust",
	}
	fmt.Println(m2)

	// access and check
	v, ok := m2["go"]
	fmt.Println(v, ok)

	// delete
	delete(m2, "rust")
	fmt.Println(m2)
}
