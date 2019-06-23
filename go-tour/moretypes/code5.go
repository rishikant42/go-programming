package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "rishi", age: 20}

	k := &p
	fmt.Println(*k)
	fmt.Println((*k).name)
	fmt.Println(k.name)

	k.name = "monu"
	fmt.Println(*k)
	fmt.Println(p)
}
