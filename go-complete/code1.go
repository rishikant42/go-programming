package main

import "fmt"
import "math/cmplx"

func main() {
	// var a int = 5
	// a := 6
	y := -5 + 12i
	var x complex128 = cmplx.Sqrt(y)
	fmt.Println("vim-go", x)
}
