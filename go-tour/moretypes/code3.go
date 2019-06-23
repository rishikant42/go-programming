package main

import "fmt"

func main() {
	i, j := 10, 15

	var p = &i
	fmt.Println("i = ", *p)
	fmt.Println("p = ", p)

	*p = 100
	fmt.Println("i = ", i)

	var t = &j
	var k = *t + 1
	fmt.Println("j = ", k)
}
