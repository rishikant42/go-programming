package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	boolValue bool
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
	pointerToPerson.boolValue = true
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 321001,
		},
	}
	jim.updateName("Jimmy")
	fmt.Printf("%+v", jim)
}
