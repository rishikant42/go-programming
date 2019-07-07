package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("world")
	wg.Wait()
}
