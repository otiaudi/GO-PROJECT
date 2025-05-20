// INSTRUCTONS
/*.
Time to order dinner!

menu is an array containing the food items on a fast food menu.

Write a loop using the range keyword with the menu array, printing each menu item and its number.

----------------------------------------------
2. 
orders maps each friend’s name to the food item they want.

Use the range keyword with the orders map to create a loop that prints each friend’s name and their order.

*/

package main

import (
  "fmt"
)

func main() {

  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

  fmt.Println("The menu:")
  // WRITE LOOP GOING THROUGH MENU IN ARRAY 
  for number, item := range menu {
    fmt.Println(number, ":", item)
  }

  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  // A late order
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friend's orders:")
  // WRITE LOOP GOING THROUGH ORDERS IN MAP
  for name, order := range orders {
    fmt.Println(name, "wants some", order)
  }

}  
