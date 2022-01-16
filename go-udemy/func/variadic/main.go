package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	fmt.Println("sum1", sum(s1...))

	s2 := sum(1, 2, 3, 4, 5)

	fmt.Println("sum2", s2)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	res := 0

	for _, v := range x {
		res += v
	}
	return res
}
