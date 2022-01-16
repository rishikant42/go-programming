package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 44

	fmt.Println("vim-go", <-ch)
}
