package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i != 5 {
			continue
		}
		fmt.Printf("Number is %d\n", i)
	}
}
