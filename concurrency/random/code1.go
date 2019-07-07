package main

import "fmt"
import "time"

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go count("Hello", c)
	go count("world", c)
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
	for msg := range c {
		fmt.Println(msg)
	}
}
