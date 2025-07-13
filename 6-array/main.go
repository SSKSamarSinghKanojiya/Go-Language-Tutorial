// Define the main package – entry point of the Go program
package main

// Import the fmt package to perform formatted I/O operations
import "fmt"

func main() {
	// Print a message to indicate the topic
	fmt.Println("We are learning array in Golang")

	// Declare a fixed-size string array of size 5
	var name [5]string

	// Assign values to specific indexes
	name[2] = "Samar"  // Assign "Samar" to index 2
	name[0] = "Vishal" // Assign "Vishal" to index 0

	// Print the entire array
	fmt.Println("Name is :", name)

	// Use %q to format array output with double quotes around each string
	fmt.Printf("Name Array is %q\n", name)

	// Uncomment the below examples to explore more array functionalities:

	/*
		// Declare and initialize an integer array with fewer elements than its size
		var numbers = [10]int{1, 2, 3, 4, 5}

		// Print the integer array
		fmt.Println("Number is :", numbers)

		// Print the length of the array using len()
		fmt.Println("Length of Number array is :", len(numbers))

		// Print the value at index 2 of the name array
		fmt.Println("Value of name at 2 index is :", name[2])
	*/
}
