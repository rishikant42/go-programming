package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	<-done
	<-done
	<-done
	<-done
	fmt.Println("elasped:", time.Since(now))
}

func task1(c chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
	c <- struct{}{}
}

func task2(c chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2")
	c <- struct{}{}
}

func task3(c chan struct{}) {
	fmt.Println("task3")
	c <- struct{}{}
}

func task4(c chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
	c <- struct{}{}
}
