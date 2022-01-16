package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go send(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
