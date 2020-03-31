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

func main() {

	// res1D := &response1{
	// 	Page:   1,
	// 	Fruits: []string{"apple", "peach", "pear"}}
	// res1B, _ := json.Marshal(res1D)
	// fmt.Println(string(res1B))

	res2D := &response2{
		XRetry: "apple",
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2D.XRetry)
	if res2D.XRetry != "" {
		fmt.Println("Retry value:", res2D.XRetry)
	} else {
		fmt.Println("Retry value is empty")
	}
	fmt.Println(string(res2B))
}
