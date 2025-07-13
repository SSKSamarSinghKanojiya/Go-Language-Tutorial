package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== File Handling Example ===")

	// -------------------------------
	// 1. Create and Write to a File
	// -------------------------------
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed at the end

	content := "Hello-World By Samar"
	bytesWritten, writeErr := io.WriteString(file, content+"\n")

	if writeErr != nil {
		fmt.Println("Error while writing to file:", writeErr)
		return
	}

	fmt.Printf("Bytes written: %d\n", bytesWritten)
	fmt.Println("File successfully created and written!")

	// -------------------------------
	// 2. Read the Entire File
	// -------------------------------
	fullContent, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	fmt.Println("\nContent using ReadFile():")
	fmt.Println(string(fullContent))

	// -------------------------------
	// 3. Read File Using Buffer
	// -------------------------------
	file2, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file2.Close()

	buffer := make([]byte, 1024)
	fmt.Println("\nContent using buffer Read():")
	for {
		n, err := file2.Read(buffer)
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}
		// Print read content
		fmt.Print(string(buffer[:n]))
	}
}












// package main

// import (
// 	"fmt"
// 	// "io"
// 	"os"
// )

// func main() {
// 	// file, err := os.Create("example.txt")
// 	// if err != nil {
// 	// 	fmt.Println("Error While creating file: ", err)
// 	// 	return
// 	// }
// 	// defer file.Close()

// 	// content := "Hello-World By Samar"
// 	// byte, errors := io.WriteString(file, content+"\n")

// 	// fmt.Println("byte written: ", byte)

// 	// if errors != nil {
// 	// 	fmt.Println("Error while writing file :", err)
// 	// 	return
// 	// }

// 	// fmt.Println("Successfully created file")

// 	/*
// 	// start
// 	// Another Learning
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Println("Error While opening file: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	// create a buffer to read the file content

// 	buffer := make([]byte, 1024)

// 	for {
// 		n, err := file.Read((buffer))
// 		if err == io.EOF { // EOF :- end of file
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error While Reading file", err)
// 			return
// 		}
// 		// Process the Read content
// 		fmt.Println(string(buffer[:n]))
// 	}
// 		// END
// 		*/

// 		// Read the entire file into a byte slice

// 		content,err := os.ReadFile("example.txt")
// 		if err != nil {
// 			fmt.Println("Error While Reading File", err)
// 			return
// 		}
// 		fmt.Println(string(content))
// }
