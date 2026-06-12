package main

import "fmt"

func main() {
	// if / else if / else
	age := 18
	if age < 18 {
		fmt.Println("minor")
	} else if age == 18 {
		fmt.Println("exactly 18")
	} else {
		fmt.Println("adult")
	}

	// if with short statement
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}

	// switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("start of week")
	case "Friday":
		fmt.Println("end of week")
	default:
		fmt.Println("midweek")
	}
}
