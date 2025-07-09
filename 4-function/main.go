package main

import "fmt"

func simpleFunction(){
	fmt.Println("Simple function")
}

func add(a,b int)(int){
	return a+b
}

func multiply(a,b int)(result int){
	result = a*b
	return
}

func main(){
	fmt.Println("We are learning function in Golang")
	simpleFunction()

	ans :=add(3,4)
	fmt.Println("add of two number is:",ans)

	data :=multiply(2,5)
	fmt.Println("Multiplication of two number is :",data)
}
