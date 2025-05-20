//An infinite loop is a loop where the condition can never be reached, causing the loop to run forever.
// In Go, you can create an infinite loop using the for statement without any condition.
// The loop will run indefinitely until it is interrupted or terminated by the program.
// In this exercise, you will create an infinite loop that repeatedly calls the count function to get the total number of guests.
package main

import (
  "fmt"
)

var number int

func count() {
  var input int
  fmt.Print("Number of guests: ")
  fmt.Scan(&input)
  number = number + input
  fmt.Println("Total guests:", number)
}

func main() {
  // WRITE THE INFINITE LOOP CONTAINING THE COUNT FUNCTION BELOW

/*  Time to count all of the members.

	   Write an infinite loop that repeatedly calls the count function to get the total number of guests!
*/

  for {
    count()
  }
}

