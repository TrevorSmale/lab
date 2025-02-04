package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func main() {
	// Prompt user for SVG file name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the SVG file (without extension): ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName) + ".svg"

	// Create a file to save the SVG
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Set canvas dimensions
	width := 500
	height := 500

	// Initialize the SVG canvas
	canvas := svg.New(file)
	canvas.Start(width, height)

	// Draw a rectangle
	canvas.Rect(50, 50, 200, 100, "fill:blue;stroke:black;stroke-width:3")

	// Draw a circle
	canvas.Circle(300, 150, 80, "fill:red;stroke:black;stroke-width:2")

	// Draw a line
	canvas.Line(100, 300, 400, 300, "stroke:black;stroke-width:5")

	// Add text
	canvas.Text(250, 400, "Go SVG Graphic!", "text-anchor:middle;font-size:24px;fill:black")

	// End and save the file
	canvas.End()

	fmt.Printf("SVG file created: %s\n", fileName)
}
