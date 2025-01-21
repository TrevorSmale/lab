package main

import (
	"fmt"
	"time"
)

func main() {
	// Location for Toronto
	torontoLoc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Loop to update and print the clock every second
	for {
		// Get current time in Toronto
		torontoTime := time.Now().In(torontoLoc)

		// Format the time with date, hour, min, second, millisecond
		clock := torontoTime.Format("2006-01-02 15:04:05")

		// Print the clock
		fmt.Println(clock)

		// Sleep for one second
		time.Sleep(time.Second)
	}
}
