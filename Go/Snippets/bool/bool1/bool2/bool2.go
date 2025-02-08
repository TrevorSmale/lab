package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Ask user for input
	fmt.Print("Enter a number between 1 and 10: ")
	var input string
	fmt.Scanln(&input)

	// Convert the input string to an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	// Define boolean conditions for checking the range
	isTooLow := number < 5
	isTooHigh := number > 5

	// Respond based on the boolean values
	if isTooLow {
		fmt.Println("Too low!")
	} else if isTooHigh {
		fmt.Println("Too high!")
	} else {
		fmt.Println("Correct! You entered 5.")
	}
}

This is a demo note change