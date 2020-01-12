package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#FF0000F",
		"green": "#DF0FFD",
	}
	colors["white"] = "#FFFFFF"
	updateMap(colors)
	printMap(colors)
}

func updateMap(c map[string]string) {
	c["red"] = "updated"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for color", color, "is", hex)
	}
}
