package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	c := Circle{5}
	s := Square{4}

	printArea(c)
	printArea(s)
}
