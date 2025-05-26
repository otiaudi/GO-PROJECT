package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup  

	wg.Add(2)  // ✅ Corrected from 'asyn.WaitGroup'
	
	
	go func(){
		time.Sleep(1 * time.Second)
		fmt.Println("Homework done!") // ✅ Changed quote style
		wg.Done()  //Tells WaitGroup this task is done
	}()
	go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("Laundry done!" )
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All chores finished!" )
}
/*
✅ Created a WaitGroup.

✅ Used wg.Add(2) to track two tasks.

✅ Launched two goroutines with go func() { ... }().

✅ Each goroutine does its job and calls wg.Done().

✅ Used wg.Wait() to wait until all chores are done.

✅ Final message is printed after both tasks complete.

✅ Used time.Sleep to simulate work being done.
*/
