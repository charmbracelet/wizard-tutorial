package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Main struct {
	index       int
	questions   []string
	answers     []string
	width       int
	height      int
	answerField textinput.Model
}

func New(questions []string) *Main {
	answerField := textinput.New()
	return &Main{questions: questions, answerField: answerField}
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
	// stack some strings together in the center of the window
	return lipgloss.JoinVertical(lipgloss.Center, m.questions[m.index], m.answerField.View())
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
