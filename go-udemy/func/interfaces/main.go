package main

import "fmt"

type Person struct {
	First string
	Last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

type Human interface {
	Speak()
}

func Bar(h Human) {
	switch h.(type) {
	case Person:
		fmt.Println("Person:", h.(Person).First)
	case SecretAgent:
		fmt.Println("ScrentAgent:", h.(SecretAgent).First)
	}
	fmt.Println("Calling human", h)
	h.Speak()
}

func (s SecretAgent) Speak() {
	fmt.Println("I am", s.Person.First, s.Last, "--From secret agent speak")
}

func (p Person) Speak() {
	fmt.Println("I am", p.First, "--From person speak")
}

func main() {
	sa1 := SecretAgent{
		Person: Person{"James", "Bond"},
		ltk:    true,
	}
	fmt.Println(sa1)
	// sa2 := SecretAgent{
	// 	Person{"Amit", "Bhadana"},
	// 	true,
	// }
	p1 := Person{"Rohit", "Sharma"}
	// sa1.Speak()
	// sa2.Speak()
	Bar(sa1)
	Bar(p1)
}
