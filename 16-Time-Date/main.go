package main

import (
	"fmt"
	"time"
)


func main(){
	currentTime := time.Now()
	fmt.Println("Current time :", currentTime)
	fmt.Printf("Type of currentTime %T\n",currentTime)

	
    // formatted := currentTime.Format("02/01/2006,Monday,15:04:05")
    formatted := currentTime.Format("02/01/2006,Monday,3:04 PM")
    fmt.Println("Formatted time:", formatted)


		// layout format
		layout_str := "25/01/2006"
		dataStr := "10/07/2025"
		formatted_time,_ := time.Parse(layout_str,dataStr)
		fmt.Println("Formatted time : ",formatted_time)

		// add 1 more day in currentTime
		new_date := currentTime.Add(24 * time.Hour)
		formatted_new_date := new_date.Format(("02/01/2006 Monday")) 
		// fmt.Println("new_date time: ", new_date.Format("02/01/2006"))
		fmt.Println("new_date time: ", formatted_new_date)
}