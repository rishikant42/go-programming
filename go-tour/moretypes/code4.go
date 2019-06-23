package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "rishi", age: 20}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)
	p.name = "monu"
	fmt.Println(p)
}
