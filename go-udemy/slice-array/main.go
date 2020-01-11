package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateStruct(pointerToPerson person) {
	pointerToPerson.firstName = "newName"
}

func updateSlice(s []string) {
	s[0] = "changed"
}

func updateArray(a [5]string) {
	a[0] = "changed"
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
	updateStruct(jim)
	jim.print()

	mySlice := []string{"Hello", "world", "this", "is", "slice"}
	updateSlice(mySlice)
	fmt.Printf("%+v\n", mySlice)

	myArray := [5]string{"Hello", "world", "this", "is", "array"}
	updateArray(myArray)
	fmt.Printf("%+v\n", myArray)
}
