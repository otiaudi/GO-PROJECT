package main

import "fmt"


func main(){
	go func(){
		fmt.Println("Hello")
	}()

	go func(){
		fmt.Println("World")
	}()

	
	  fmt.Println("Fin")
		
	select{}
}

