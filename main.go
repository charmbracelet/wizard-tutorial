package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Main struct {
	index     int
	questions []string
	answers   []string
	width     int
	height    int
}

func New(questions []string) *Main {
	return &Main{questions: questions}
}

func (m Main) Init() tea.Cmd {
	return nil
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Main) View() string {
	if m.width == 0 {
		return "loading..."
	}
	return "got width"
}

func main() {
	// start bubble tea and init first model
	questions := []string{"what is your name?", "what is your favourite editor?", "what language do you like most?"}
	firstQuestion := New(questions)
	p := tea.NewProgram(*firstQuestion, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
