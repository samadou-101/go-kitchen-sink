package main

import "fmt"

func main() {
	// Integers
	// Signed
	//   int8, int16, int32(rune), int64, int (32 or 64 bit)
	// Unsinged
	//   uint8(byte), uint16, uint32, uint64, uint
	// Special
	//    uintptr (stores memory address)

	// explicit type
	var name string = "name"

	// type inferece
	var age = 101

	// short declaration
	country := "CountryName"

	// multiple vars

	var x, y = 10, 20

	// zero values
	var score int
	var active bool
	var message string

	fmt.Println(name, age, country)
	fmt.Println(x, y)
	fmt.Println(score, active, message)

}
