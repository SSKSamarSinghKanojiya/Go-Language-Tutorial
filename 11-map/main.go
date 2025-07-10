package main

import "fmt"

func main(){
	// name <-> grade 
	studentGrades := make(map[string]int)

	studentGrades["Samar"] = 100
	studentGrades["Vishal"] = 90
	studentGrades["Poonam"] = 85
	studentGrades["Laxmi"] = 80

	fmt.Println("Marks of Vishal :", studentGrades["Vishal"])
	studentGrades["Vishal"] = 150
	fmt.Println("New Marks of Vishal :", studentGrades["Vishal"])

	delete(studentGrades,"Vishal")
	fmt.Println("Marks of Vishal :", studentGrades["Vishal"])

	// Checking if key exists
	fmt.Println()
	Grades, Exists := studentGrades["Samar"]
	fmt.Println("Grades of Samar :",Grades)
	fmt.Println("Samar exists :",Exists)

	for index, value := range studentGrades{
		fmt.Printf("Key is %s and marks is %d\n", index,value)
	}

	fmt.Println()
	person :=map[string]int{
		"Father Sahab" : 500,
		"MAA": 501,
	}

	for index, value := range person {
		fmt.Printf("Key is %s and marks is %d\n", index,value)
	}
} 