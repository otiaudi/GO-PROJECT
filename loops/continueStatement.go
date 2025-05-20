package main

import "fmt"


func main(){
	JellyBeans := []string {"green", "blue", "yellow", "red", "green", "yellow", "red"}

	for index :=0; index<len(JellyBean) index++{
		if JellyBean[index] == "green" {
			continue
		}
		fmt.Println("You ate the", JellyBean[index], "JellyBean")
	}
}
