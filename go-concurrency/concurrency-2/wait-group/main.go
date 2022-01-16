package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go work(&wg, i+1)
	}
	wg.Wait()
	fmt.Println("Elapsed:", time.Since(now))
}

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
	// wg.Done()
}
