package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello")
	say("world")
	// time.Sleep(4000 * time.Millisecond)
}
