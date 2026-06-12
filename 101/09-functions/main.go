package main

import "fmt"

// simple function
func add(a, b int) int {
	return a + b
}

// multiple return values
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// named return values
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

// variadic
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(div(10, 2))
	fmt.Println(split(10))
	fmt.Println(sum(1, 2, 3, 4))
}
