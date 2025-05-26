package main

import (
	"fmt"
	"time"
)
// The 'pinger' function waits for a signal from the 'pinger' channel,
// prints "ping", then sends a signal to the 'ponger' channel.
func pinger(pinger <-chan int, ponger chan<- int){
	for{
		<-pinger        // Wait to receive a signal from 				 //the 'pinger' channel
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <-1    // Send signal to 'ponger' channel 			      to continue
	} 
}

// The ponger print a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int){
	for{
		<-ponger
		fmt.Println("Pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main(){
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

// the main goroutine starts the ping/pong by sending into the ping channel
ping <-1

for{
   // block the main thread until an interrupt
   time.Sleep(time.Second) // Keep the main function running (prevents exit)

   }
}
