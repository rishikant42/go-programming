package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	XRetry string   `json:"x_retry"`
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type response3 []struct {
	x int
}

func main() {
	var x response3
	type y struct {
		x int
	}
	z := y{x: 1}
	x = append(x, z)
	fmt.Println(x)
	var res []response1

	res1D := response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res = append(res, res1D)
	fmt.Println(res)

	res2D := &response2{
		XRetry: "apple",
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
