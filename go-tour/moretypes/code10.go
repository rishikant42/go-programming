package main

import "fmt"

func main() {
	var m = make(map[string]string)
	m["a"] = "apple"
	m["b"] = "boy"
	fmt.Println(m)
	m["a"] = "ant"
	fmt.Println(m)
	fmt.Println(m["aa"])
	delete(m, "b")
	fmt.Println(m)

	d := make(map[int]int)
	d[0] = 100
	fmt.Println(d[0])
}
