package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, UOR! \n")
	})

	// /hello endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, NIBM! \nThis is endpoint 2\n")
	})



	fmt.Println("Starting UOR API Server...")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}