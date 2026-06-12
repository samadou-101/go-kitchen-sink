package main

import "fmt"

// generic function with type parameter
func Last[T any](s []T) T {
	return s[len(s)-1]
}

// comparable constraint
func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Last([]int{1, 2, 3}))
	fmt.Println(Last([]string{"a", "b", "c"}))
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"a", "b"}, "c"))
}
