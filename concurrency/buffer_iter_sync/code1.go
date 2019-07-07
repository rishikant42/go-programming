package main

import "fmt"
import "sync"

var wg sync.WaitGroup
var mut sync.Mutex

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
	mut.Lock()
	defer mut.Unlock()
	fmt.Println("Inside goroutine")
}

func main() {
	fooCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooCh, i)
	}
	wg.Wait()
	close(fooCh)
	for item := range fooCh {
		fmt.Println(item)
	}
	// for {
	// 	fooValue, open := <-fooCh
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(fooValue)
	// }
}
