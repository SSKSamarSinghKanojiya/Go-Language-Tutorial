// Define the main package, which is the entry point of a Go program.
package main

// Import required packages:
// - "bufio" for buffered input (helps read full lines including spaces).
// - "fmt" for formatted I/O (e.g., printing).
// - "os" to access the operating system, like reading from standard input.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Ask the user for their name.
	fmt.Println("Hey, What's your name?")

	// Create a new reader to read input from the user (from standard input - the keyboard).
	reader := bufio.NewReader(os.Stdin)

	// Read the entire input line until the newline character (`\n`).
	// `name` will store the input, and `_` ignores the error value returned by ReadString.
	name, _ := reader.ReadString('\n')

	// Print a greeting message using the user's input.
	// %s is used to format the string, and `name` is inserted in place of %s.
	fmt.Printf("Hello, %s", name)

	// fmt.Println("Hey, What's your name?")
	// var name string
	// fmt.Scan(&name) // This only captures input up to the first space
	// fmt.Println("Hello, Mr.", name)

}
