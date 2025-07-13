package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("=== URL Parsing and Modification Example ===")

	// A full URL string
	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", myUrl)

	// Parse the URL string
	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Can't parse URL:", err)
		return
	}
	fmt.Printf("Type after parsing: %T\n", parsedURL)

	// Access individual components
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)

	// -----------------------------
	// Modify Path and Query Params
	// -----------------------------
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamsamar"

	// Reconstruct the new URL
	newUrl := parsedURL.String()
	fmt.Println("New Modified URL:", newUrl)
}
