package main

import ("fmt"
	"time"
)

func task(id int){
	fmt.Printf("Task %d is starting\n", id)
	 time.Sleep(2 * time.Second)
	 fmt.Printf("Task %d is completed\n", id)
 }


 func main(){
	 go task(1)

	 go task(2)

	 time.Sleep(5*time.Second) //waiting for both task to complete
	 fmt.Println("All task competed concurrently")
 }

	
