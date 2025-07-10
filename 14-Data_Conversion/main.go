package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	num = num + 8
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	data = data + 5.50
	fmt.Println("Data is ", data)
	fmt.Printf("Type of Data is %T\n", data)

	// Conversion number into string
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Str is ", str)
	fmt.Printf("Type of Str is %T\n", str)

	// Conversion string into number
	number_string := "100"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 50
	fmt.Println("number_int is ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	// Conversion string into float 
	num_string := "3.14"
	number_float,_ := strconv.ParseFloat(num_string,64)
	fmt.Println("number_float is ",number_float)
	fmt.Printf("Type of number_float is %T\n",number_float)
	 
}
