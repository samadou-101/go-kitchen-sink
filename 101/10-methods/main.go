package main

import "fmt"

type Rect struct {
	W, H float64
}

// value receiver
func (r Rect) Area() float64 {
	return r.W * r.H
}

// pointer receiver (can modify)
func (r *Rect) Scale(factor float64) {
	r.W *= factor
	r.H *= factor
}

func main() {
	r := Rect{3, 4}
	fmt.Println(r.Area())

	r.Scale(2)
	fmt.Println(r)
}
