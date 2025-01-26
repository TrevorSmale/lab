package main

import (
	"fmt"
	"time"
)

func worker1() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Worker 1 is running")
	}
}

func worker2() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Worker 2 is running")
	}
}

func worker3() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Worker 3 is running")
	}
}

func main() {
	// Start each function as a goroutine
	go worker1()
	go worker2()
	go worker3()

	// Keep the main program running for 3 seconds to observe the output
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Main program finished")
}
