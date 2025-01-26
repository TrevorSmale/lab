package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a reader to get user input
	reader := bufio.NewReader(os.Stdin)

	// Ask the user if the sun is up
	fmt.Print("Is the sun up? (yes/no): ")
	userInput, _ := reader.ReadString('\n')

	// Clean up the input (remove newline and make lowercase)
	userInput = strings.TrimSpace(strings.ToLower(userInput))

	// Determine if the sun is up based on input
	var sunIsUp bool
	if userInput == "yes" {
		sunIsUp = true
	} else if userInput == "no" {
		sunIsUp = false
	} else {
		fmt.Println("Invalid input! Please enter 'yes' or 'no'.")
		return
	}

	// Decide whether to turn the light on or off
	if sunIsUp {
		fmt.Println("The sun is up! Turning the light OFF.")
	} else {
		fmt.Println("The sun is down! Turning the light ON.")
	}

	// Bonus: Show the light's status as a boolean
	lightIsOn := !sunIsUp
	fmt.Println("Is the light on?", lightIsOn)
}
