package main

import (
	"fmt"
	"time"
	"os"
)
// select block is used to wait on multiple channel operations
// It is like switch case for channels
// It allows a goroutine to wait on multiple communication operations
// If multiple cases are ready, one case is chosen at random
// If no case is ready, the select statement blocks until one case is ready

func Select(c chan int, quits chan struct{}){
// switch case does for control statement
// select does for channel async
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Waiting for channel to be ready")
		// select block
		select{
		case <-c:
			fmt.Println("Recieved")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0) // exit the program
		}
	}
						

}

func main(){  
// fmt.Println("vim-go")
c:= make(chan int)
quits:= make(chan struct{})

go func(){
	c <-1
	quits <- struct{}{} // send empty struct to quits channel
	time.Sleep(2 * time.Second)
}()
go Select(c, quits)

select {}
}
