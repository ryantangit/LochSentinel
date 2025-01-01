package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


func main() {
	
	publicDir := "./public/"
	publicFiles := http.FileServer(http.Dir(publicDir))
	http.Handle("/", publicFiles)

	http.HandleFunc("/api/ping", pingResponseHandler)
	
	fmt.Println("Listening on port 6009")
	err := http.ListenAndServe(":6009", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func pingResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct for the response
	type pingResponse struct {
		Message string    `json:"message"`
	}

	// Create a dynamic response
	response := pingResponse{
		Message: "Pong",
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}


