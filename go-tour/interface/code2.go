package main

import (
	"fmt"
)

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I passed into bar", h)
	// h.speak()

	switch h.(type) {
	case person:
		fmt.Println("first person", h.(person).first)
	case secretAgent:
		fmt.Println("first secretAgent", h.(secretAgent).first)
	}
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	sa2 := secretAgent{
		person: person{"Miss", "MoneyPenny"},
		ltk:    true,
	}
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	// sa1.speak()
	// sa2.speak()
	// fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 5
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int = int(x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
}
