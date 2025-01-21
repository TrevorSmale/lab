package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Question struct represents a question with follow-up questions for choices "a" and "b"
type Question struct {
	MainQuestion string
	FollowUpA    string
	FollowUpB    string
}

// Model represents the application state
type Model struct {
	Question      Question
	CurrentChoice string
}

// Init initializes the application state
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles events and updates the application state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			return m, tea.Quit
		case "b":
			return m, tea.Quit
		case "esc", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the UI
func (m Model) View() string {
	return fmt.Sprintf("Main Question: %s\n\nChoose 'a' for: %s\nChoose 'b' for: %s\n", m.Question.MainQuestion, m.Question.FollowUpA, m.Question.FollowUpB)
}

func main() {
	// Define the initial question
	initialQuestion := Question{
		MainQuestion: "Do you prefer apples or bananas?",
		FollowUpA:    "Why do you prefer apples?",
		FollowUpB:    "Why do you prefer bananas?",
	}

	// Create a new Model with the initial question
	initialModel := Model{
		Question:      initialQuestion,
		CurrentChoice: "",
	}

	// Start the Bubbletea program
	program := tea.NewProgram(initialModel)
	if err := program.Start(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
