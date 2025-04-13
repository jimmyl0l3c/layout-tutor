package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimmyl0l3c/layout-tutor/courses"
	"github.com/jimmyl0l3c/layout-tutor/layout"
	"github.com/jimmyl0l3c/layout-tutor/menu"
)

type model struct {
	view int

	courseMenu menu.Model
	levelMenu  menu.Model
}

func initialModel() model {
	// ti := inputfield.New()
	// ti.TextToWrite = "sons seas tree stories inns"
	// ti.Focus()
	// ti.Width = 10
	return model{
		courseMenu: menu.New("Choose course", []list.Item{courses.Colemak}),
		levelMenu:  menu.New("Choose level", []list.Item{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.courseMenu.SetSize(msg.Width, msg.Height)
		m.levelMenu.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "-":
			switch m.view {
			case 1:
				m.view = 0
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if c, ok := m.courseMenu.GetSelected().(layout.LayoutCourse); ok {
				items := make([]list.Item, len(c.Levels))

				for i, v := range c.Levels {
					items[i] = v
				}

				cmd := m.levelMenu.SetItems(items)
				m.view = 1

				return m, cmd
			}
		}
	}

	var cmd tea.Cmd

	switch m.view {
	case 0:
		m.courseMenu, cmd = m.courseMenu.Update(msg)
		return m, cmd
	case 1:
		m.levelMenu, cmd = m.levelMenu.Update(msg)
		return m, cmd
	}

	// m.textInput, cmd = m.textInput.Update(msg)

	return m, nil
}

func (m model) View() string {
	// msg = m.textInput.View()

	switch m.view {
	case 0:
		return m.courseMenu.View()
	case 1:
		return m.levelMenu.View()
	}

	return "Unknown view"
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something failed: %v", err)
		os.Exit(1)
	}
}
