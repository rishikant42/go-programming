package main

import (
	"fmt"
)

var mapping = map[string]bool{
	"apple":    true,
	"boy":      true,
	"cat":      true,
	"dog":      false,
	"elephant": true,
	"horse":    false,
}

func getMap() string {
	return "apple"
}

func testMapping() (string, bool) {
	switch getMap() {

	case "apple":
		for k, v := range mapping {
			if !v {
				return k, v
			}
		}
	default:
		return "bye", false
	}
	return "hello", false
}

func main() {
	a, b := testMapping()
	fmt.Println(a, b)
}
