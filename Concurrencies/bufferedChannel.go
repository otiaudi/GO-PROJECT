package main

import(
	"fmt"

)
// Buffered channels allow you to send multiple values without waiting for them to be received.
// They can be useful when you want to decouple the sending and receiving of values.
// Buffered channels can be created using the make function with a second argument specifying the buffer size.
// Note that if the buffer is full, sending to the channel will block until space is available.
// If the buffer is empty, receiving from the channel will block until a value is sent.
// Buffered channels can be used to improve performance in certain scenarios, such as when sending multiple values in a loop.
func  main(){
	c:=make(chan int, 3)

	func(){
		c<-1
		c<-2
		c<-3 
		close(c)
	}()
// Receiving from a buffered channel can be done using a for loop.
	for i:= range c{
		fmt.Println(i)

	}
}
