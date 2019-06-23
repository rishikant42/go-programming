package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [3]int{1, 2}
	fmt.Println(primes)

	var abc [2]int
	abc[1] = 3
	fmt.Println(abc)
}
