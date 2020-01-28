package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area: ", s.getArea())
}

func main() {
	t := triangle{height: 4, base: 5}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}
