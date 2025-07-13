// Define the main package — the entry point for the Go program.
package main

// Import the "fmt" package for formatted I/O operations and error formatting.
import "fmt"

// divide is a function that takes two float64 numbers as input.
// It returns two values:
// - the result of division (float64)
// - an error (if any)
// If the denominator (b) is zero, it returns an error.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Return 0 and a formatted error if division by zero is attempted.
		return 0, fmt.Errorf("denominator must not be zero")
	}
	// Otherwise, return the result and nil for error.
	return a / b, nil
}

func main() {
	fmt.Println("Started error handling in functions")

	// Call the divide function with 10 and 0.
	ans, err := divide(10, 0)

	// Check if an error occurred during division
	if err != nil {
		// Print the error message and exit the function early.
		fmt.Println("Error:", err)
		return
	}

	// If no error, print the division result.
	fmt.Println("Division of two numbers is:", ans)
}
