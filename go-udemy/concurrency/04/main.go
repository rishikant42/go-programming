package main

import "fmt"

func main() {
	ch := make(<-chan int, 2)

	fmt.Println("------")
	fmt.Printf("%T\n", ch)
}
