package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, "as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Its boolean")
		case int, int64:
			fmt.Println("Its integer")
		case string:
			fmt.Println("Its string")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	var num1 int64 = 4
	var num2 int32 = 4
	whatAmI(true)
	whatAmI("Hey")
	whatAmI(1)
	whatAmI(num1)
	whatAmI(num2)
	fmt.Println("vim-go")
}
