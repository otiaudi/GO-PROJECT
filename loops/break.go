package main

import "fmt"


func main(){
	// Break statement
	animal:=[]string{"Cat", "Dog", "Fish", "Turtle"}

	for index:=0; index<len(animal); index++ {
		if animal[index] == "Dog"{
			fmt.Println("This the perfect Animal")
			break // stop searching the array
		}
		
	}
	// Continue statement
	jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
	for index := 0; index < len(jellybeans); index++ {
		if jellybeans[index] == "green" {
			continue
  }
  fmt.Println("You ate the", jellybeans[index], "jellybean!")
}

}
