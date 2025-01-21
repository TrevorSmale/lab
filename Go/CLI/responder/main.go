package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A command-line application",
	Long:  `A simple command-line application written in Go using Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display a usage message when no subcommand is provided
		fmt.Println("Please provide a subcommand. Use --help for more information.")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet a person by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s!\n", name)
	},
}

var goodbyeCmd = &cobra.Command{
	Use:   "goodbye [name]",
	Short: "Bid farewell to a person by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Goodbye, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(goodbyeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

