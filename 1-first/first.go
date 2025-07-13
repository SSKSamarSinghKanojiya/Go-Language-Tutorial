// Declare the package name.
// This is used to group related Go files together.
// The package name here is "first", which means this file belongs to the "first" package.
package first

// Import the "fmt" package from the Go standard library.
// "fmt" stands for "format", and it provides functions to format text, including printing to the console.
import "fmt"

// PrintMessage is a function that takes a single parameter named "message" of type string.
// The purpose of this function is to print the message to the console.
func PrintMessage(message string) {
	// Use the fmt.Println function to print the message to the console.
	// fmt.Println automatically adds a newline at the end of the message.
	fmt.Println(message)
}
