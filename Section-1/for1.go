// running inifinite loop
// This program will run an infinite loop until the user enters "exit".
// It will then ask the user if they want to exit, and if they confirm, it will break the loop and exit the program.
// If the user enters anything other than "exit", it will continue to prompt for a name.
package main
import (
	"fmt"
) 

func main() {
	for{
		var name string
		fmt.Println("Enter your name:")
		fmt.Scanln(&name)
		if name == "exit"{
			fmt.Println("Exiting the program.")
			var exit string
			fmt.Println("Do you want to exit? (yes/no)")
			fmt.Scanln(&exit)
			if exit == "yes"{
				fmt.Println("Exiting the program.")

			break
		}
	}
	}
}

