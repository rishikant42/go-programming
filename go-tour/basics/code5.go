package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Sum is ", add(2, 3))
}
