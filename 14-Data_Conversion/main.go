package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Type Conversion Example ===")

	// ----------------------------
	// int → float64
	// ----------------------------
	var num int = 42
	num = num + 8
	fmt.Println("Number is:", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num) // Explicit conversion
	data = data + 5.50
	fmt.Println("Data is:", data)
	fmt.Printf("Type of data is %T\n", data)

	// ----------------------------
	// int → string using strconv.Itoa()
	// ----------------------------
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Str is:", str)
	fmt.Printf("Type of str is %T\n", str)

	// ----------------------------
	// string → int using strconv.Atoi()
	// ----------------------------
	number_string := "100"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 50
	fmt.Println("number_int is:", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	// ----------------------------
	// string → float64 using strconv.ParseFloat()
	// ----------------------------
	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is:", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}
