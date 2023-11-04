package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Folks! \n")
	})

	// /hello endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Folks! \nThis is endpoint 2\n")
	})



	fmt.Println("Starting Folks API Server...")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
