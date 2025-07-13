package main

import "fmt"

// Function that returns the sum of two integers
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")

	// Call add() and print result
	data := add(5, 6)
	fmt.Println("Data is", data)

	// Defer statement - executed at the end of the function (just before main() exits)
	defer fmt.Println("Middle of the program")

	// This will execute before the deferred statement
	fmt.Println("End of the program")
}
