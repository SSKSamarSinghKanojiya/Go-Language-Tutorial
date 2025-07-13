// Package main is the entry point of any Go program.
package main

import (
	"fmt"
)

// =====================
// Function Section
// =====================

// divide performs division and handles divide-by-zero errors.
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
	if x > 5 && (y > 5 || z < 43) {
		fmt.Println("Hey, How are you?")
	} else {
		fmt.Println("Learn programming with Hello World")
	}

	// --------------------------------------
	// Part 5: Switch Statements
	// --------------------------------------
	fmt.Println("\n=== Switch Statements Example ===")

	// Basic switch with single values
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	// Switch with multiple case values
	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	case "July", "August", "September":
		fmt.Println("Monsoon")
	default:
		fmt.Println("Other season")
	}

	// Switch without expression (like if-else)
	temperature := -4
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
