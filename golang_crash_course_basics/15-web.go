package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About stuff</h1>")
}

func main() {
	fmt.Println("Server running on http://localhost:5000")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":5000", nil)
}
