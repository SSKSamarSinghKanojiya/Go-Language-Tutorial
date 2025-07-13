// Define the main package.
// This is the starting point of a Go application.
package main

// Import the "fmt" package which provides functions for formatted I/O operations like printing to the console.
import "fmt"

// The main function is the entry point of any Go program.
func main() {
	// Declare and initialize a variable "age" of type int with value 25
	age := 25

	// Declare and initialize a variable "name" of type string with value "Alice"
	name := "Alice"

	// Declare and initialize a variable "height" of type float64 with value 5.108
	height := 5.108

	// Print multiple variables using fmt.Println.
	// This will automatically insert spaces and a newline at the end.
	fmt.Println("age:", age, "height:", height, "name:", name)

	// Print a static string message
	fmt.Println("Hello World")

	// Print using formatted string with fmt.Printf
	// %d is used for integers
	fmt.Printf("age is %d\n", age)

	// %.3f is used to print float values with 3 digits after the decimal point
	fmt.Printf("height is %.3f\n", height)

	// %T prints the data type of the variable
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of height is %T\n", height)

	// %s is used to print string values
	fmt.Printf("name is %s\n", name)
}
