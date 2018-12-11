package main

import (
    "fmt"
    "math"
    "math/rand"
)


func sqrt(n float64) float64 {
    return math.Sqrt(n)
}

func main() {
    fmt.Println("The square root of 4 is ", sqrt(4))
    fmt.Println("Random number in b/w 1-100 is ", rand.Intn(4))
}
