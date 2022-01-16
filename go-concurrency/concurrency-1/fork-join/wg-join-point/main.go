package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()
	fmt.Println("done waiting, Main exits")
	fmt.Println("Elasped:", time.Since(now))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done work")
}
