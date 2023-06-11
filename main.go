package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, NIBM! \nHow you like that? \n")
	})

	fmt.Println("Starting NIBM API Server...")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}