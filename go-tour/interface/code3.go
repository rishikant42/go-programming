package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a, b Abser
	f := MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	b = &v
	fmt.Println(b.Abs())
}
