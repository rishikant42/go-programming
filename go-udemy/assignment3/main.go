package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: filename arg is missing")
		os.Exit(1)
	}
	fname := os.Args[1]
	fopen, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error:", err)
	}
	io.Copy(os.Stdout, fopen)
}
