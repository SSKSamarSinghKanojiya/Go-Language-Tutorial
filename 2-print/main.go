package main

import "fmt"

func main(){
	age := 25
	name := "Alice"
	height := 5.108

	fmt.Println("age:" ,age, "height:" ,height, "name:", name)
	fmt.Println("Hello World")

	fmt.Printf("age is %d\n",age)
	fmt.Printf("heigth is %.3f\n",height)
	fmt.Printf("Type of age is %T\n",age)
	fmt.Printf("Type of name is %T\n",name)
	fmt.Printf("Type of height is %T\n",height)

	fmt.Printf("name is %s\n",name)
}