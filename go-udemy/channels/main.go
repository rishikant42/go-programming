package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, ch chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}
