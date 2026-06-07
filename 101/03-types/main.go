package main

import "fmt"

/*  - Basic types: int, uint, float32, bool, string
	- Alias types: byte(uint8), rune(int32)
    - Composite types: array, slice, map, struct
    - Reference types: pointer, function, interface, channel
    - Special: nil(zero value for poitners, slices, maps, interfaces, channels, functions)
*/

func main() {
	var num byte = 255
	var r rune = 'A'

	// array
	arr := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}

	m := map[string]int{"a": 1, "b": 2}

	type User struct {
		Name string
		Age  int
	}

	u := User{"Ali", 25}

	fmt.Println(num, r)
	fmt.Println(arr, slice, m, u)
}
