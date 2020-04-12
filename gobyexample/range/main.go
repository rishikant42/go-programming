package main

import "fmt"

func main() {
	array := [3]string{"aa", "bb", "cc"}
	for index, value := range array {
		fmt.Println("index:", index, "value:", value)
	}

	slice := []string{"apple", "boy", "cat"}
	for index, value := range slice {
		fmt.Println("index:", index, "value:", value)
	}

	mapping := map[string]string{"a": "apple", "b": "boy", "c": "cat"}
	for key, value := range mapping {
		fmt.Println("key:", key, "value:", value)
	}
	for key := range mapping {
		fmt.Println("key:", key)
	}

	for i, c := range "golan" {
		fmt.Println("i:", i, "c", c)
	}
}
