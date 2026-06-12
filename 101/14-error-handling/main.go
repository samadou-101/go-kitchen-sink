package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// idiomatic error check
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("error:", err2)
	} else {
		fmt.Println(result2)
	}
}
