package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.length * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rect) other() float64 {
	return r.length + r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var s Shape
	fmt.Println("Value of s is", s)
	fmt.Printf("Type of s is %T\n", s)

	s = Rect{3, 4}
	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("Rect Area value is %v\n", s.Area())
	fmt.Printf("Rect Perimeter value is %v\n", s.Perimeter())

	var r Rect
	r = Rect{3, 4}
	fmt.Println("s == r is", s == r)
	fmt.Printf("Type of r is %T\n", s)
	fmt.Printf("Rect Area value is %v\n", r.Area())
	fmt.Printf("Rect Perimeter value is %v\n", r.Perimeter())
	fmt.Printf("Rect Perimeter value is %v\n", r.other())

	s = Circle{3}
	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("Rect Area value is %v\n", s.Area())
	fmt.Printf("Rect Perimeter value is %v\n", s.Perimeter())
}
