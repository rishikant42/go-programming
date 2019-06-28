package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
	Area() float64
}

type Cube struct {
	Side float64
}

func (c Cube) Area() float64 {
	return 6 * c.Side * c.Side
}

func (c Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}

func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("Area = ", s.Area())
	fmt.Println("Volume = ", o.Volume())
	fmt.Println("Area = ", c.Area())
	fmt.Println("Volume = ", c.Volume())
}
