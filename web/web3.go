package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>Hello there</p>")
	fmt.Fprintf(w, "<p>Hello %s, I am %s</p>", "john", "<h2>rishi</h2>")

	fmt.Fprintf(w, `<h1>Hey there</h1>
<p>Hello there</p>
<p>Hello there</p>
<p>Hello %s, I am %s</p>`, "john", "<h2>rishi</h2>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8888", nil)
}
