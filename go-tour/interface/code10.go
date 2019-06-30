package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Metrial interface {
	Shape
	Object
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
	var m Metrial = c
	var s Shape = c
	var o Object = c
	fmt.Println("Area = ", s.Area())
	fmt.Println("Volume = ", o.Volume())
	fmt.Println("Volume = ", m.Area())
	fmt.Println("Volume = ", m.Volume())
}
