package main

import "fmt"


func main(){
	letter := [] string{"A", "B", "C", "D"}

	for index, value := range letter{
		fmt.Println("index:", index, "value :", value)
	}
}
