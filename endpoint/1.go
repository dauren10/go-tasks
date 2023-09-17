package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write the response content
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a new HTTP server
	server := http.NewServeMux()

	// Define a route that maps the /hello endpoint to the helloHandler function
	server.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":2020", server)
}
