package main

import (
	"fmt"
)

type Vertex struct {
	x, y float64
}

func (v *Vertex) ScaleMethod(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFn(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScaleMethod(10)
	fmt.Println(v)

	(&v).ScaleMethod(10)
	fmt.Println(v)

	p := &v
	p.ScaleMethod(10)
	fmt.Println(v)

	ScaleFn(&v, 10)
	fmt.Println(v)
}
