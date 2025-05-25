package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Run printNumbers in a goroutine

	time.Sleep(3 * time.Second) // Give the goroutine time to finish
	fmt.Println("Main function completed.")
}
// In this example, the `printNumbers`
//  function is run in a goroutine,
//  allowing it to execute concurrently
//  with the main function.
//  The main function waits
//  for a few seconds before printing its
//  completion message.
// Goroutines are lightweight threads
//  managed by the Go runtime. 
// They are used to perform tasks 
// concurrently, allowing for efficient 
// use of system resources and improved 
// performance in applications that require
//  parallel processing.