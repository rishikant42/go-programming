package main

import "fmt"
import "math"
import "strings"

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("String type", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("Int type", i)
	case float64:
		fmt.Println("float64 type", math.Sqrt(i.(float64)))
	default:
		fmt.Println("Unknown type", i)
	}
}

func main() {
	s := "Hello world"
	i := 5
	f := 4.0
	b := true
	explain(s)
	explain(i)
	explain(f)
	explain(b)
}
