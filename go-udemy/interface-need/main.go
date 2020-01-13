package main

import "fmt"

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

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {
	// eb := englishBot{}
	sb := spanishBot{}
	// printGreeting(eb)
	printGreeting(sb)
}
