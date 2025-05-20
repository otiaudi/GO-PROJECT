package main


import "fmt"


func main(){
	 addressBook := map[string] string{
		 "John": "12 Main St",
		 "Janet": "56 Pleasant St",
		 "Jordan": "88 Liberty Ln",
	 }

	 for key, value := range addressBook {
		 fmt.Println("Name :", key, "Address :", value)
	 }
 }
