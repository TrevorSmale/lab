// Dup1 prints the text of each line that appears mor than once in the stdio, preceded by count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("dup.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close() // Ensure the file is closed when done

	// Create a map to count occurrences
	counts := make(map[string]int)

	// Create a scanner to read the file line by line
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}

	// Check for scanner errors
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}

	// Print lines that appear more than once
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
