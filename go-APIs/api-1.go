package main

import (
	"fmt"
	"net/http"
)

func main() {
	//This creates a new ServeMux, which is a request multiplexer
	// A ServeMux matches the URL of each incoming request against a list of registered patterns
	// and calls the handler for the pattern that most closely matches the URL.
	
	mux := http.NewServeMux() // Create a new ServeMux
	// Register a handler function for the "/great" route
	// The handler function writes "Hello world" to the response writer
	// The handler function is called whenever a request is made to the "/great" route.
	// The function also prints "Request Received" to the console.
	// The response writer is used to send the response back to the client.f
	
	mux.HandleFunc("/great", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Received")
		w.Write([]byte("Hello world\n"))
	})
	
	// added another route called "Hello"
	mux.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Second Request Recieved")
		w.Write([]byte("This is my second Request"))
	})

	// Start the server on localhost:3001
	// ListenAndServe listens on the TCP network address 
	// and then calls Serve to handle requests on incoming connections.
	// If the server fails to start, it will print an error message.

	err := http.ListenAndServe("localhost:3001", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
