package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println("m:", m)
	fmt.Println("m[one]:", m["one"])
	delete(m, "one")
	fmt.Println("m:", m)

	v, ok := m["twoo"]
	fmt.Println("m[two] value:", v)
	fmt.Println("ok:", ok)

	m2 := map[string]string{"a": "apple"}
	fmt.Println("m2:", m2)
}
