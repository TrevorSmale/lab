package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Get current date
	currentDate := time.Now()

	// Format date in YY.MM format
	folderName := fmt.Sprintf("%02d.%02d", currentDate.Year()%100, currentDate.Month())

	// Create new directory
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create 4 folders within the new directory
	for i := 1; i <= 4; i++ {
		folderPath := fmt.Sprintf("./%s/folder%d", folderName, i)
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}

	fmt.Printf("Directory %s created successfully with 4 folders.\n", folderName)
}
