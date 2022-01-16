package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work()
		done <- struct{}{}
	}()

	fmt.Println("Channel msg", <-done)
	fmt.Println("done waiting, Main exits")
	fmt.Println("Elasped:", time.Since(now))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done work")
}
