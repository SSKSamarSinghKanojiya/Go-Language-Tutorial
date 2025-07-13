// Define the main package. This is the entry point of the Go program.
package main

// Import the "fmt" package to allow formatted I/O operations like printing.
import "fmt"

// simpleFunction is a basic function with no parameters and no return value.
// It just prints a message to the console.
func simpleFunction() {
	fmt.Println("Simple function")
}

// add is a function that takes two integers `a` and `b` as input,
// and returns a single integer (their sum).
func add(a, b int) int {
	return a + b
}

// multiply is a function that takes two integers `a` and `b`,
// and returns their product. It uses named return value `result`.
func multiply(a, b int) (result int) {
	result = a * b // assign the result
	return         // returns `result` implicitly
}

// main function is the entry point of the Go program.
func main() {
	fmt.Println("We are learning function in Golang")

	// Call simpleFunction (no arguments, just prints)
	simpleFunction()

	// Call add function with 3 and 4, store result in `ans`
	ans := add(3, 4)
	fmt.Println("Addition of two numbers is:", ans)

	// Call multiply function with 2 and 5, store result in `data`
	data := multiply(2, 5)
	fmt.Println("Multiplication of two numbers is:", data)
}
