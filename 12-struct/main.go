package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct{
	Email string
	Phone string
	Fax string
}

type Address struct{
	House int
	Area string
	State string
}

type Employee struct{
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var samar Person
	fmt.Println("Samar Person :", samar)
	samar.FirstName = "Samar"
	samar.LastName = "Singh"
	samar.Age = 24

	// fmt.Println("Samar Person : ", samar)

	// 2nd Methods
	person1 := Person{
		FirstName: "Vishal",
		LastName: "Singh",
		Age: 28,
	}
	fmt.Println("Person1 : ",person1)

	// New Keyword Method
	var person2 = new(Person)
	person2.FirstName = "Raju"
	person2.LastName = "Boss"
	person2.Age = 25

	// fmt.Println("Person2 : ",person2)

	var employee1 Employee
	// employee1 := Employee //Same
	employee1.Person_Details = Person{
		FirstName: "Ritik",
		LastName: "Kumawat",
		Age: 26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "987654321"
	employee1.Person_Contact.Fax = "Fax@4567890"

	employee1.Person_Address = Address{
		House: 12,
		Area: "Indore",
		State: "Madhya Pradesh",
	}

	fmt.Println("Employee 1: ",employee1.Person_Details)
	fmt.Println("Employee 1: ",employee1.Person_Contact)
	fmt.Println("Employee 1: ",employee1.Person_Address)
}
