package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	length, width float64
}

func (r Rect) Area() float64 {
	return r.length * r.width
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var s Shape
	r := Rect{3, 4}
	s = &r
	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("Rect Area value is %v\n", s.Area())
	fmt.Printf("Rect Perimeter value is %v\n", s.Perimeter())
}
