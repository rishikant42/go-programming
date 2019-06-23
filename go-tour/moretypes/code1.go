package main

import "fmt"

func main() {
	i, j := 10, 20
	fmt.Println(i)
	fmt.Println(j)

	p := &i
	fmt.Println(p)
	*p += 2
	fmt.Println(*p)
	fmt.Println(i)
	// for ; *p < 24; *p++ {
	// 	fmt.Println(*p)
	// 	fmt.Println(i)
	// }
}
