package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func main() {
	// Prompt for SVG file name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the SVG file (without extension): ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName) + ".svg"

	// Create file for SVG
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Set canvas dimensions
	width := 1200
	height := 800

	// Init canvas
	canvas := svg.New(file)
	canvas.Start(width, height)

	// Draw rect
	canvas.Rect(0, 0, 1200, 800, "fill:green;stroke:black;stroke-width:0")

	// Draw line
	canvas.Line(0, 300, 1200, 300, "stroke:white;stroke-width:4")

	// Add text
	canvas.Text(250, 400, "Certificate of Completion", "text-anchor:middle;font-size:24px;fill:black")

	// End and save the file
	canvas.End()

	fmt.Printf("SVG file created: %s\n", fileName)
}
