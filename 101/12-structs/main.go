package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// struct literal
	p1 := Person{"Alice", 30}
	fmt.Println(p1)

	// named fields
	p2 := Person{Name: "Bob", Age: 25}
	fmt.Println(p2.Name)

	// zero value struct
	var p3 Person
	p3.Name = "Charlie"
	fmt.Println(p3)

	// pointer to struct
	p4 := &Person{Name: "Dave", Age: 40}
	fmt.Println(p4.Age)
}
