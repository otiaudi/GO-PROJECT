package main

import (
	"fmt"
	"time"
)

// Function that runs an endless loop, printing "Heavy" every 1 second
func heavy() {
	for {
		time.Sleep(time.Second * 1) // wait for 1 second
		fmt.Println("Heavy")
	}
}

// Function that runs an endless loop, printing "Super Heavy" every 2 seconds
func superHeavy() {
	for {
		time.Sleep(time.Second * 2) // wait for 2 seconds
		fmt.Println("Super Heavy")
	}
}

func main() {
	// Start both functions in separate goroutines (background workers)
	go heavy()       // Starts printing "Heavy" every 1 second in background
	go superHeavy()  // Starts printing "Super Heavy" every 2 seconds in background

	fmt.Println("Fin") // This prints once, right after launching the goroutines

	// select{} keeps the main function running forever
	// Without it, the program would exit immediately and stop the goroutines
	select {}
}
