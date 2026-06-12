package main

import "fmt"

func main() {
	// standard for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// for as while
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop (commented to avoid hanging)
	// for {}

	// for range over slice
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// for range over map
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
