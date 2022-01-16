package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go work(&wg)

	wg.Wait()
	fmt.Println("done waiting, Main exits")
}

func work(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done work")
	wg.Done()
}
