package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 44
	ch <- 40

	fmt.Println("vim-go", <-ch)
	fmt.Println("vim-go", <-ch)
}
