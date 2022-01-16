package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 44
	}()

	fmt.Println("vim-go", <-ch)
}
