package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("=== HTTP GET Request Example ===")

	// Make a GET request to a public API
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close() // Ensure response body is closed

	// Print the type of the response (pointer to http.Response)
	fmt.Printf("Type of response: %T\n", res)

	// Read the response body into bytes
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Convert byte slice to string to view JSON content
	fmt.Println("Response body:", string(data))
}
