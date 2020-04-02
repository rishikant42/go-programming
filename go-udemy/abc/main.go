package main

import "fmt"

func main() {
	abc := "hello"
	var xyz string
	if abc == "hello" {
		xyz = "wow"
		fmt.Println("vim-go")
	} else {
		fmt.Println("vim-bye")
	}
	fmt.Println("xyz=", xyz)
}
