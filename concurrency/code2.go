package main

import "fmt"

func sum(s int, n int, c chan int) {
	result := 0
	for i := s; i < n; i++ {
		result += i
	}
	c <- result
}

func main() {
	c := make(chan int)
	go sum(1, 3333333333, c)
	go sum(1, 3333333333, c)
	go sum(1, 3333333333, c)
	// sum2 := sum(11, 21)
	// fmt.Println(sum1 + sum2)
	sum1, sum2 := <-c, <-c
	fmt.Println(sum1 + sum2)
}
