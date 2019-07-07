package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
	wg.Done()
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("This is panic")
		}
		fmt.Println(i, s)
	}
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("world")
	wg.Wait()
}
