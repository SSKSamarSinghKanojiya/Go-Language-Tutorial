package main

import "fmt"

func main(){
	for i := 1; i <= 10; i++ {
		fmt.Println("Number is :",i)
	}

	counter := 0

	for{
		fmt.Println("Infinite Lopp",counter)
		counter++
		if  counter == 3 {
			break
		}
	}

	numbers := []int{11,22,33,44,55}
	for index,value := range numbers{
		fmt.Printf("Index of Number is %d, and value is %d\n",index,value)
	}

	data := "Hello, World"
	for index, char := range data{
		fmt.Printf("Index of Data is %d, and value is %c\n",index,char)
	}
}