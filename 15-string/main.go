package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split
	data := "Apple,Orange,Banana,Mango,Fruits"
	parts := strings.Split(data, ",,")
	fmt.Println(parts)


	// Count
	str := "one two three four two two five"
	count := strings.Count(str,"two")
	fmt.Println("count of two is :", count)

	// TrimSpace
	str = "    Hello, Go!      "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed: ",trimmed)

	str1 := "Samar"
	str2 := "Singh"
	result := strings.Join([]string{str1,str2}," ")
	fmt.Println("Result : ", result)
}
