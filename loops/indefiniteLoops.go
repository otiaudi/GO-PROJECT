package main

import "fmt"


func ask()(int){
	fmt.Println("I am thinking of a number")
	// get the input from the user
	var input int
	fmt.Scan(&input)

	return input
}


func main(){
	//call the function ask and store its value in a variable

	var guess int
	for guess != 56 {
		guess = ask()
	}
	
	if guess == 56 {
		fmt.Println("Your guess perfect, you have wo $2000, congrats")
	}
}
