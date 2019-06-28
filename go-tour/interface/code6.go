package main

import "fmt"

type MyString string

type Rect struct {
	length float64
	width  float64
}

func explain(i interface{}) {
	fmt.Printf("Type: %T\nValue: %v\n", i, i)
}

func main() {
	ms := MyString("Hello world")
	r := Rect{3, 4}
	explain(ms)
	explain(r)
}
