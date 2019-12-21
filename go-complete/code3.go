package main

import "fmt"

func inc(i *int) {
	*i++
}

func main() {
	var ap *int
	i := 10
	ap = &i
	inc(ap)
	fmt.Println(*ap)
}
