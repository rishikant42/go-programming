package main

import (
	"fmt"
)

func add(a int, b int) (c int) {
	c = a + b
	return
}

func main() {
	fmt.Println(add(2, 1))
}
