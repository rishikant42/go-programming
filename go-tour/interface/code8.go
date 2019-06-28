package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Skin interface {
	Color() float64
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
	var s1 Shape = Cube{3}
	var c1, ok1 = s1.(Cube)
	fmt.Printf("value = %v\nok = %v\n", c1, ok1)
	fmt.Println("Area = ", c1.Area())
	fmt.Println("Volume = ", c1.Volume())

	var s2 Shape = Cube{3}
	var c2, ok2 = s2.(Skin)
	fmt.Printf("value = %v\nok = %v\n", c2, ok2)

	var s3 Object = Cube{3}
	var c3, ok3 = s3.(Cube)
	fmt.Printf("value = %v\nok = %v\n", c3, ok3)
	fmt.Println("Area = ", c3.Area())
	fmt.Println("Volume = ", c3.Volume())

}
