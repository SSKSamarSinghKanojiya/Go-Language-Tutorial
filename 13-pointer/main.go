package main

import "fmt"

// Function that modifies the value using a pointer
func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	fmt.Println("=== Pointer Example ===")

	// -----------------------------
	// Basic Pointer Declaration
	// -----------------------------
	num := 2
	ptr := &num // ptr holds the address of num

	// Print original values
	fmt.Println("Num has value:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Pointer contains:", ptr)
	fmt.Println("Value via pointer (dereference):", *ptr)

	fmt.Println()

	// -----------------------------
	// Nil Pointer Check
	// -----------------------------
	var pointer *int // default is nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned (nil)")
	}

	fmt.Println()

	// -----------------------------
	// Pass by Reference Example
	// -----------------------------
	value := 5
	modifyValueByReference(&value) // Pass address of value to function
	fmt.Println("Value after modifyValueByReference():", value)
}
