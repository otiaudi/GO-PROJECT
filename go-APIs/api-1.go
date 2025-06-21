package main

import (
	"fmt"
	"net/http"
)

func main() {
	//used to create a new multiplexer 
	//that routes incoming HTTP requests 
	// to the correct handler functions.
	mux := http.NewServeMux()

	// Register a handler function for the "/great" path.
	mux.HandleFunc("/great", func(w http.ResponseWriter, r *http.Request) {
		// Log that the request was received.
		fmt.Println("Request Received")
		// Write a plain-text response back to the client.
		w.Write([]byte("Hello world\n"))
	})

	// Register another handler for the "/Hello" path.
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Println("Second Request Received")
		// Write a response back to the client.
		w.Write([]byte("This is my second Request"))
	})

	// Start the HTTP server on localhost port 3001.
	// ListenAndServe takes in an address and a handler (our mux).
	err := http.ListenAndServe("localhost:3001", mux)
	if err != nil {
		// If the server fails to start, log the error.
		fmt.Println("Server error:", err)
	}
}
/* This code sets up a simple HTTP server 
   with two endpoints:"/great" and "/Hello".
  
   When a request is made to either 
   endpoint, it responds with a message 
   and logs the request.
  
   The server listens on localhost at port 
   3001.To test this, you can run the code 
   and then use a web browser or a tool 
   like curl to access the endpoints:as
     - For the "/great" endpoint, navigate 
       to http://localhost:3001/great
     - For the "/Hello" endpoint, navigate 
	   to http://localhost:3001/Hello
 */