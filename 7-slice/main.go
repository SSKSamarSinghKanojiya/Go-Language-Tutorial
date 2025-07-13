// Package main is the entry point of any Go application.
package main

import (
	"fmt"
)

// ====================
// Function Section
// ====================

// divide performs division of two float64 numbers and handles divide-by-zero errors.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

// ====================
// Main Function
// ====================

func main() {

	// --------------------------------------
	// Part 1: Working with Arrays
	// --------------------------------------
	fmt.Println("=== Array Example ===")

	var name [5]string
	name[2] = "Samar"
	name[0] = "Vishal"

	fmt.Println("Name array content:", name)
	fmt.Printf("Name Array (quoted): %q\n", name)

	/*
		// Uncomment for integer array example
		var numbers = [10]int{1, 2, 3, 4, 5}
		fmt.Println("Number array:", numbers)
		fmt.Println("Length of number array:", len(numbers))
		fmt.Println("Value at index 2 of name array:", name[2])
	*/

	// --------------------------------------
	// Part 2: Function with Error Handling
	// --------------------------------------
	fmt.Println("\n=== Function Error Handling Example ===")

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred during division:", err)
	} else {
		fmt.Println("Division result:", ans)
	}

	// Try with valid input
	ans2, err2 := divide(20, 4)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Division of 20 by 4 is:", ans2)
	}

	// --------------------------------------
	// Part 3: Working with Slices
	// --------------------------------------
	fmt.Println("\n=== Slice Example ===")

	// Create a slice of int using make: len = 3, cap = 5
	numbers := make([]int, 3, 5)
	// Append additional values to the slice
	numbers = append(numbers, 4)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))     // Current number of elements
	fmt.Println("Capacity:", cap(numbers))   // Total allocated space

	// Create an empty slice with length 0
	stock := make([]int, 0)
	fmt.Println("Stock slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
}
