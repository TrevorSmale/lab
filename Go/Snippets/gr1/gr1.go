package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from goroutine!")
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func main() {
	// Start the sayHello function as a goroutine
	go sayHello()

	// Main function keeps running
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from main!")
		time.Sleep(500 * time.Millisecond)
	}

	// Wait for a moment to ensure the program doesn't exit too soon
	time.Sleep(2 * time.Second)
}
