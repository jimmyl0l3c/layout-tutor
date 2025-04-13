package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jimmyl0l3c/layout-tutor/layout"
	"github.com/jimmyl0l3c/layout-tutor/layout/courses"
	"github.com/jimmyl0l3c/layout-tutor/views"
	"github.com/jimmyl0l3c/layout-tutor/views/course"
	"github.com/jimmyl0l3c/layout-tutor/views/menu"
)

type model struct {
	view views.View

	courseMenu menu.Model
	levelMenu  menu.Model
	courseView course.Model
}

func initialModel() model {
	return model{
		courseMenu: menu.New("Choose course", []list.Item{courses.Colemak}),
		levelMenu:  menu.New("Choose level", []list.Item{}),
		courseView: course.New(),
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
		m.courseView.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "-":
			switch m.view {
			case views.LevelMenu:
				m.view = views.LayoutMenu
			case views.Course:
				m.view = views.LevelMenu
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.view {
			case views.LayoutMenu:
				if c, ok := m.courseMenu.GetSelected().(layout.LayoutCourse); ok {
					m.courseView.SetLayout(c)

					items := make([]list.Item, len(c.Levels))
					for i, v := range c.Levels {
						items[i] = v
					}

					cmd := m.levelMenu.SetItems(items)
					m.view = views.LevelMenu

					return m, cmd
				}
			case views.LevelMenu:
				if l, ok := m.levelMenu.GetSelected().(layout.Level); ok {
					m.courseView.SetLevel(l)
					m.view = views.Course

					return m, nil
				}
			}
		}
	}

	var cmd tea.Cmd

	switch m.view {
	case views.LayoutMenu:
		m.courseMenu, cmd = m.courseMenu.Update(msg)
		return m, cmd
	case views.LevelMenu:
		m.levelMenu, cmd = m.levelMenu.Update(msg)
		return m, cmd
	case views.Course:
		m.courseView, cmd = m.courseView.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.view {
	case views.LayoutMenu:
		return m.courseMenu.View()
	case views.LevelMenu:
		return m.levelMenu.View()
	case views.Course:
		return m.courseView.View()
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
