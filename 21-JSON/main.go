
package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct with JSON tags for correct key mapping
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("=== JSON Encoding and Decoding ===")

	// Create a person object
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("Original Person struct:", person)

	// ----------------------------
	// Marshal: Go Struct → JSON
	// ----------------------------
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("JSON Output:", string(jsonData))

	// ----------------------------
	// Unmarshal: JSON → Go Struct
	// ----------------------------
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("Decoded Person struct:", personData)
}













// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	IsAdult bool   `json:"is_adult"`
// }

// func main() {
// 	fmt.Println("We are learning JSON")
// 	person := Person{Name: "John", Age: 34, IsAdult: true}
// 	fmt.Println("Person Data is :", person)

// 	// Convert person into JSON Encoding (Marshalling)
// 	jsonData, err := json.Marshal(person)

// 	if err != nil {
// 		fmt.Println("Error marshalling", err)
// 		return
// 	}

// 	fmt.Println("Person Data is : ", string(jsonData))


// 	//  Decoding (Unmarshalling)
// 	var personData Person
// 	err = json.Unmarshal(jsonData, &personData)

// 	if err != nil {
// 		fmt.Println("Error unmarshalling ",err)
// 		return
// 	}
// 	fmt.Println("person Data is after unmarshalling: ",personData)
// }
