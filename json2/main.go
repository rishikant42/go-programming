package main

import (
	"fmt"
)

type Base struct {
	Length string `json:"length,omitempty"`
	Width  string `json:"width,omitempty"`
}

type Dummy struct {
	Base
	Name    string  `json:"name,omitempty"`
	Number  int64   `json:"number,omitempty"`
	Pointer *string `json:"pointer,omitempty"`
	Errorrr error
}

type newDummy struct {
	NewName string
	Age     int64
}

func (d *Dummy) createNewDummy() (*newDummy, error) {
	nd := &newDummy{}
	nd.NewName = d.Name
	nd.Age = d.Number
	return nd, nil
}

func main() {
	pointer := "yes"

	dummyComplete := Dummy{}
	dummyComplete.Length = "5"
	dummyComplete.Number = 10
	dummyComplete.Pointer = &pointer

	newdummy, _ := dummyComplete.createNewDummy()
	fmt.Printf("%+v\n", dummyComplete)
	fmt.Printf("%+v\n", *newdummy)
	fmt.Println("name=", dummyComplete.Name)
	fmt.Println("new name=", newdummy.NewName)

	// fmt.Println("base:", dummyComplete.Length)
	// fmt.Printf("%+v", dummyComplete)

	// if dummyComplete.Errorrr != nil {
	// 	fmt.Println("error:", dummyComplete.Errorrr)
	// } else {
	// 	fmt.Println("No error")
	// }

	// data, err := json.Marshal(dummyComplete)
	// if err != nil {
	// 	fmt.Println("An error occured: ", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(data))

	// // // ALL of those are considered empty by Go
	// // dummyNilPointer := Dummy{
	// // 	Name:    "",
	// // 	Number:  1,
	// // 	Pointer: nil,
	// // }

	// // dataNil, err := json.Marshal(dummyNilPointer)
	// // if err != nil {
	// // 	fmt.Println("An error occured: ", err)
	// // 	os.Exit(1)
	// // }

	// // fmt.Println(string(dataNil))
}
