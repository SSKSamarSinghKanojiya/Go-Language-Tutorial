package main

import "fmt"


func main(){
	fmt.Println("We are learning array in Golang")

	// var name[5]string
	// name[0]="Samar";
	// name[2]="Raju";

	// fmt.Println("Name of Person is :", name)

	// var numbers = [10]int{1,2,3,4,5}
	// fmt.Println("Number is :",numbers)
	// fmt.Println("Length of Number array is :",len(numbers))
	// fmt.Println("Value of name at 2 index is :", name[2])

	// ""

	var name[5]string
	name[2] = "Samar"
	name[0] = "Vishal"

	fmt.Println("Name is :",name)
	fmt.Printf("Name Array is %q\n",name)
}
