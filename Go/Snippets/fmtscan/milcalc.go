package main

import "fmt"

func main() {
	var inches float64 // Use float64 for better precision
	var millimeters float64

	// Prompt the user to enter a number of inches
	fmt.Print("Enter the number of inches: ")
	fmt.Scan(&inches)

	// Convert inches to millimeters
	millimeters = inches * 25.4

	// Display the result
	fmt.Printf("%.2f inches is equal to %.2f millimeters.\n", inches, millimeters)
}
