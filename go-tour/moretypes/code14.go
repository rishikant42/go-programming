package main

import (
	"fmt"
)

func add(x float64, y float64) int {
	return int(x + y)
}

func main() {
	x := float64(10)
	y := 20.2
	fmt.Println(x + y)
}
