package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"goapis/structs" // 👈 This matches your module name in go.mod
)

func main() {
	// Create multiplexer
	mux := http.NewServeMux()

	// Register /ping route
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//call structs.Response 
			//(correct package reference)
			data := structs.Respose{
				Code: http.StatusOK,
				 Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	fmt.Println("Server running at http://localhost:3002/ping")
	http.ListenAndServe(":3002", mux)
}
