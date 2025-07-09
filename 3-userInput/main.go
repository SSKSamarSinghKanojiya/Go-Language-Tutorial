package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// fmt.Println("Hey, What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.",name)

	fmt.Println("Hey, What's your name?")
	reader := bufio.NewReader(os.Stdin)
	name,_:= reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)

}
 