package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application's state
type model struct {
	progress int    // Current progress (0-100)
	maxWidth int    // Maximum width of the progress bar
	emoji    string // The emoji to represent the progress bar
	quitting bool   // Quit flag
}

// Initialize the model
func initialModel() model {
	return model{
		progress: 0,
		maxWidth: 100, // The bar will have 20 blocks max
		emoji:    "⏰", // Emoji representing progress
		quitting: false,
	}
}

// TickMsg signals a timer tick
type TickMsg struct{}

// The Init function starts the timer
func (m model) Init() tea.Cmd {
	// Start ticking every 100ms
	return tickCmd()
}

// Command to send a TickMsg after 100ms
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

// Update processes messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q": // Quit the program
			m.quitting = true
			return m, tea.Quit
		}

	case TickMsg:
		if m.progress < 100 {
			m.progress++
			return m, tickCmd() // Continue ticking until complete
		}
	}

	return m, nil
}

// View renders the UI
func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	// Calculate the current bar length
	barLength := (m.progress * m.maxWidth) / 100
	bar := strings.Repeat(m.emoji, barLength)

	return fmt.Sprintf(
		"Progress: [%s%s] %d%%\n\nPress 'q' to quit.\n",
		bar, strings.Repeat(" ", m.maxWidth-barLength), m.progress,
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}
}
