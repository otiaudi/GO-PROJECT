// This tells Go that this is the main package – like the starting point of your program
package main

// Here, we're bringing in some useful tools (called packages) from Go
import (
	"fmt"      // For printing messages to the screen
	"net/http" // For building a web server and handling browser requests
)

// This is the function that Go looks for to start running your code
func main() {
	// This tells the server: "When someone goes to the home page (/), use the homePage function to respond"
	http.HandleFunc("/", homePage)

	// Print a message to the terminal so we know the server is running
	fmt.Println("Server is running on port 8080")

	// Start the web server on port 8080. The 'nil' means we are using the default settings.
	http.ListenAndServe(":8080", nil)
}

// This is the function that handles what the user sees when they visit the home page
func homePage(w http.ResponseWriter, r *http.Request) {
	// We use Fprint to send a message back to the browser
	
	fmt.Fprint(w, "Welcome to the Book API!")
}

