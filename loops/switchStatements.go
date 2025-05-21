package main
import (
	"fmt"
)

func main() {
	var day string
	fmt.Println("Enter  your fevorite days")
	fmt.Scanln(&day)

	switch day {
	case "Mon","Tue", "Wed":
		fmt.Println("Working days")

	case "Fri","Sat","sun":
		fmt.Println(" Are Weekends")

	case "fri","sat":
		fmt.Println("Fun days")

	default:
		fmt.Println("Thanks for trying")
	}
}

