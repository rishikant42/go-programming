package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello")
	go say("world")
	time.Sleep(time.Millisecond * 1000)
}
