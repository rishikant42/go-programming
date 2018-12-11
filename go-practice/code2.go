package main

import "fmt"


func add(x, y float64) float64 {
    return x+y
}

func main() {
    a, b := 4.6, 6.3
    fmt.Println("Sum =", add(a, b))
}
