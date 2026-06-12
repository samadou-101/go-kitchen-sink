package main

import "fmt"

func main() {
	// slice declaration (nil)
	var s []int
	fmt.Println(s, len(s), cap(s))

	// make
	s2 := make([]int, 3, 5)
	s2[0] = 1
	fmt.Println(s2, len(s2), cap(s2))

	// slice literal
	s3 := []string{"a", "b", "c"}
	fmt.Println(s3)

	// append
	s3 = append(s3, "d")
	fmt.Println(s3)

	// slicing
	sub := s3[1:3]
	fmt.Println(sub)
}
