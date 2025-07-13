package main

import "fmt"

func main() {
	// --------------------------------------
	// Part 7: Maps
	// --------------------------------------
	fmt.Println("=== Map Example ===")

	// Create an empty map: name -> grade
	studentGrades := make(map[string]int)

	// Add entries
	studentGrades["Samar"] = 100
	studentGrades["Vishal"] = 90
	studentGrades["Poonam"] = 85
	studentGrades["Laxmi"] = 80

	// Accessing a value by key
	fmt.Println("Marks of Vishal:", studentGrades["Vishal"])

	// Updating a value
	studentGrades["Vishal"] = 150
	fmt.Println("New Marks of Vishal:", studentGrades["Vishal"])

	// Deleting a key
	delete(studentGrades, "Vishal")
	fmt.Println("Marks of Vishal after deletion:", studentGrades["Vishal"]) // Will show 0 (zero value)

	// Check if a key exists
	fmt.Println()
	grade, exists := studentGrades["Samar"]
	fmt.Println("Grades of Samar:", grade)
	fmt.Println("Samar exists:", exists)

	// Loop through the map
	for name, mark := range studentGrades {
		fmt.Printf("Student: %s, Marks: %d\n", name, mark)
	}

	// Initialize a map with values directly
	fmt.Println()
	person := map[string]int{
		"Father Sahab": 500,
		"MAA":          501,
	}

	for relation, value := range person {
		fmt.Printf("Relation: %s, Value: %d\n", relation, value)
	}
}
