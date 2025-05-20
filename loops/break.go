package main

import "fmt"


func main(){
	animal:=[]string{"Cat", "Dog", "Fish", "Turtle"}

	for index:=0; index<len(animal); index++ {
		if animal[index] == "Dog"{
			fmt.Println("This the perfect Animal")
			break // stop searching the array
		}
		
	}
}
