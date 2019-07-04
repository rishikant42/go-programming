package main

import "fmt"

func sum(s []int, c chan int) {
	result := 0
	for _, v := range s {
		result += v
	}

	c <- result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	a := s[:len(s)/2]
	b := s[len(s)/2:]
	go sum(a, c)
	go sum(b, c)

	x, y := <-c, <-c
	fmt.Println(a, b)
	fmt.Println(x, y, x+y)
}
