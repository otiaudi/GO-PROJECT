package main

import "fmt"

func main() {
	// creating robot channels
	robot1 := make(chan string)
	robot2 := make(chan string)

	// create go function for robot1
	go func() {
		robot1 <- "🧹 Sweeping done!"
	}()

	// create go function for robot2
	go func() {
		robot2 <- "🧼 Washing done!"
	}()

	// receive messages from channels
	receiver1 := <-robot1
	receiver2 := <-robot2

	fmt.Println(receiver1)
	fmt.Println(receiver2)
}
