package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Listening on port 6009")
	err := http.ListenAndServe(":6009", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
