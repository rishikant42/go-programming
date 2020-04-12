package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Array a:", a)

	a[3] = 3
	fmt.Println("Array a:", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b:", b)

	var c [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("Array c:", c)
}
