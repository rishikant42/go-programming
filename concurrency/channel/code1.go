package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooCh := make(chan int)
	go foo(fooCh, 4)
	go foo(fooCh, 5)
	v1 := <-fooCh
	v2 := <-fooCh
	fmt.Println(v1, v2)
}
