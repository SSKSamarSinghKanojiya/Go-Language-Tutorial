package main

import "fmt"

func main(){
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// 	// break

	// case 2:
	// 	fmt.Println("Tuesday")

	// case 3:
	// 	fmt.Println("Wednesday")
	
	// default:
	// 	fmt.Println("Unknown day")
	// }

	// Second Switch case
	// month := "January"
	// month := "1"

	// switch month {
	// case "January", "February", "March":
	// 	fmt.Println("Winter")
	// case "April","May","June":
	// 	fmt.Println("Spring")
	// case "1":
	// 	fmt.Println("Good Day we are very fine")
	// default:
	// 	fmt.Println("Other season")
	// }

temperature := -4

	switch{
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >=10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 20:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}