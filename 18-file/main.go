package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error While creating file: ", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Hello-World By Samar"
	// byte, errors := io.WriteString(file, content+"\n")

	// fmt.Println("byte written: ", byte)

	// if errors != nil {
	// 	fmt.Println("Error while writing file :", err)
	// 	return
	// }

	// fmt.Println("Successfully created file")

	/*
	// start
	// Another Learning
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error While opening file: ", err)
		return
	}
	defer file.Close()

	// create a buffer to read the file content

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read((buffer))
		if err == io.EOF { // EOF :- end of file
			break
		}
		if err != nil {
			fmt.Println("Error While Reading file", err)
			return
		}
		// Process the Read content
		fmt.Println(string(buffer[:n]))
	}
		// END
		*/

		// Read the entire file into a byte slice

		content,err := os.ReadFile("example.txt")
		if err != nil {
			fmt.Println("Error While Reading File", err)
			return
		}
		fmt.Println(string(content))
}
