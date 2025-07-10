package main

import "fmt"

func modifyValueByReference(num *int){
	*num = *num + 20
}


func main(){
	// One way define
	// var num int 
	// var ptr *int 
	// ptr = &num

	// Another way define
	num := 2
	ptr := &num
	fmt.Println("Num has value : ", num)
	fmt.Println("Num has value : ", &num)
	fmt.Println("Pointer contains : ",ptr)
	fmt.Println("Data contains through Pointer : ",*ptr)

	fmt.Println()
	// By default pointer value is nil
	var pointer *int 
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value :=5
	modifyValueByReference(&value)
	fmt.Println("Value contains : ",value)
}