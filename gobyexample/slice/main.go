package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "apple"
	s[1] = "boy"
	s[2] = "cat"
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))

	s = append(s, "dog")
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))

	s = append(s, "elephant", "fish")
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)
	fmt.Println("len(c):", len(c))

	fmt.Println("s[2:4]:", s[2:4])
	fmt.Println("s[1]:", s[1])
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[:4]:", s[:4])

	d := []int{1, 1}
	fmt.Println("d:", d)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD:", twoD)
}
