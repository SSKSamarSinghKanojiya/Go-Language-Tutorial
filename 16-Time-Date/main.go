package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Time and Date Example ===")

	// ----------------------------
	// Get current time
	// ----------------------------
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime)

	// Format current time in a readable format
	formatted := currentTime.Format("02/01/2006, Monday, 3:04 PM")
	fmt.Println("Formatted current time:", formatted)

	// ----------------------------
	// Parse a string into a time
	// ----------------------------
	layoutStr := "25/01/2006" // Must match the layout format
	dateStr := "10/07/2025"
	parsedTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Parsed time from string:", parsedTime)

	// ----------------------------
	// Add 1 day to the current time
	// ----------------------------
	newDate := currentTime.Add(24 * time.Hour)
	formattedNewDate := newDate.Format("02/01/2006 Monday")
	fmt.Println("Date after 1 day:", formattedNewDate)
}
