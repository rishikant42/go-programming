package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println("j = ", j)
	}

	for {
		fmt.Println("Infinte loop")
		break
	}

	for k := 1; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("k = ", k)
	}
	fmt.Println("vim-go")
}
