package main

import "fmt"

func sliceContains(s []string, e string) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	s := []string{"apple", "boy", "cat", "dog"}
	fmt.Println("vim-go:", sliceContains(s, "aaaaaa"))
}
