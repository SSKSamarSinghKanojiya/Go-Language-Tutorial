package main

import "fmt"

func main() {
	x := 21
	if x > 5 {
		fmt.Printf("%d is greater than 5\n",x)
	} else {
		fmt.Printf("%d is smaller than 5\n",x)
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	}else{
		fmt.Println("z is smaller than 5")
	}

	y := 10
	if x > 5 && (y > 5 || z < 43) {
		fmt.Println("Hey, How are you ?")
	} else {
		fmt.Println("Learn programming with Hello World")
	}
}
