package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimmyl0l3c/layout-tutor/inputfield"
)

type model struct {
	textInput inputfield.Model
}

func initialModel() model {
	ti := inputfield.New()
	ti.TextToWrite = "sons seas tree stories inns"
	ti.Focus()
	ti.Width = 10

	return model{
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return inputfield.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	msg := "Type what you see\n\n"

	msg += m.textInput.View()

	return msg
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something failed: %v", err)
		os.Exit(1)
	}
}
