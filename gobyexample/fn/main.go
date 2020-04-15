package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) int {
	fmt.Println("nums:", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("simple example")
	fmt.Println("1+2=", plus(1, 2))
	fmt.Println("1+2+3=", plusPlus(1, 2, 3))

	fmt.Println("Multiple return value")
	a, b := vals()
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	_, c := vals()
	fmt.Println("c=", c)

	fmt.Println("Variadic example")
	x := sum(1, 2, 3, 4)
	fmt.Println("x:", x)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	y := sum(s...)
	fmt.Println("y:", y)

	fmt.Println("closures example")

	fmt.Println("intSeq()():", intSeq()())
	fmt.Println("intSeq()():", intSeq()())

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())

	fmt.Println("Recursion example")

	fmt.Println("fact(7):", fact(7))
}
