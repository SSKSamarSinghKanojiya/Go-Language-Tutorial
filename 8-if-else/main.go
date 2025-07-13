// Package main is the entry point of the Go program.
package main

import (
	"fmt"
)

// =====================
// Function Section
// =====================

// divide performs division of two float64 numbers and returns an error if the denominator is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

// =====================
// Main Function
// =====================

func main() {

	// --------------------------------------
	// Part 1: Arrays
	// --------------------------------------
	fmt.Println("=== Array Example ===")

	var name [5]string
	name[2] = "Samar"
	name[0] = "Vishal"

	fmt.Println("Name array content:", name)
	fmt.Printf("Name Array (quoted): %q\n", name)

	/*
		// Uncomment to explore integer array
		var numbers = [10]int{1, 2, 3, 4, 5}
		fmt.Println("Number array:", numbers)
		fmt.Println("Length of number array:", len(numbers))
		fmt.Println("Value at index 2 of name array:", name[2])
	*/

	// --------------------------------------
	// Part 2: Error Handling in Function
	// --------------------------------------
	fmt.Println("\n=== Function Error Handling Example ===")

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred during division:", err)
	} else {
		fmt.Println("Division result:", ans)
	}

	ans2, err2 := divide(20, 4)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Division of 20 by 4 is:", ans2)
	}

	// --------------------------------------
	// Part 3: Slices
	// --------------------------------------
	fmt.Println("\n=== Slice Example ===")

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 4)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	stock := make([]int, 0)
	fmt.Println("Stock slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))

	// --------------------------------------
	// Part 4: Conditional Statements
	// --------------------------------------
	fmt.Println("\n=== Conditional Statements Example ===")

	x := 21
	if x > 5 {
		fmt.Printf("%d is greater than 5\n", x)
	} else {
		fmt.Printf("%d is smaller than 5\n", x)
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than or equal to 10")
	} else {
		fmt.Println("z is smaller than or equal to 5")
	}

	y := 10
	// Using logical AND (&&) and OR (||) operators
	if x > 5 && (y > 5 || z < 43) {
		fmt.Println("Hey, How are you?")
	} else {
		fmt.Println("Learn programming with Hello World")
	}
}
