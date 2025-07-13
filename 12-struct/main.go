package main

import "fmt"

// Define a basic struct for person details
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Define a struct for contact details
type Contact struct {
	Email string
	Phone string
	Fax   string
}

// Define a struct for address details
type Address struct {
	House int
	Area  string
	State string
}

// Define a composite struct combining Person, Contact, and Address
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	fmt.Println("=== Structs Example ===")

	// Method 1: Default declaration, then set fields
	var samar Person
	fmt.Println("Samar Person (initial):", samar)
	samar.FirstName = "Samar"
	samar.LastName = "Singh"
	samar.Age = 24

	// Method 2: Inline struct initialization
	person1 := Person{
		FirstName: "Vishal",
		LastName:  "Singh",
		Age:       28,
	}
	fmt.Println("Person1:", person1)

	// Method 3: Using new() keyword
	var person2 = new(Person)
	person2.FirstName = "Raju"
	person2.LastName = "Boss"
	person2.Age = 25
	fmt.Println("Person2:", *person2) // dereference to print value

	// Composite Struct: Employee with embedded structs
	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Ritik",
		LastName:  "Kumawat",
		Age:       26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "987654321"
	employee1.Person_Contact.Fax = "Fax@4567890"
	employee1.Person_Address = Address{
		House: 12,
		Area:  "Indore",
		State: "Madhya Pradesh",
	}

	// Print nested struct values
	fmt.Println("Employee Details:", employee1.Person_Details)
	fmt.Println("Employee Contact:", employee1.Person_Contact)
	fmt.Println("Employee Address:", employee1.Person_Address)
}
