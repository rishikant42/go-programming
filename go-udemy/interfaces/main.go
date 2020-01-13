package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// some custom logic specific to englishBot
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// some custom logic specific to spanishBot
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
