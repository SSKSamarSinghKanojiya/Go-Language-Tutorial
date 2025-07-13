package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== String Operations Example ===")

	// ----------------------------
	// strings.Split() - Break a string into parts using a separator
	// ----------------------------
	data := "Apple,Orange,Banana,Mango,Fruits"
	parts := strings.Split(data, ",") // Use "," instead of ",," here
	fmt.Println("Split result:", parts)

	// ----------------------------
	// strings.Count() - Count how many times a substring appears
	// ----------------------------
	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("Count of 'two' is:", count)

	// ----------------------------
	// strings.TrimSpace() - Remove leading and trailing spaces
	// ----------------------------
	str = "    Hello, Go!      "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed string:", trimmed)

	// ----------------------------
	// strings.Join() - Join strings using a separator
	// ----------------------------
	str1 := "Samar"
	str2 := "Singh"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Joined string:", result)
}
