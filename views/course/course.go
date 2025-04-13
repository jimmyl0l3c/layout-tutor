package course

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/layout-tutor/inputfield"
	"github.com/jimmyl0l3c/layout-tutor/layout"
)

type Model struct {
	layout layout.LayoutCourse
	level  layout.Level

	contentStyle lipgloss.Style
	titleStyle   lipgloss.Style
	inputStyle   lipgloss.Style

	input inputfield.Model
	help  help.Model
}

func New() Model {
	m := Model{
		contentStyle: lipgloss.NewStyle().Margin(1, 2),
		titleStyle: lipgloss.NewStyle().
			MarginBottom(1).
			Padding(0, 1).
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")),
		inputStyle: lipgloss.NewStyle().MarginBottom(1),

		input: inputfield.New(),
		help:  help.New(),
	}
	m.input.Focus()
	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Reset):
			m.input.Reset()
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	title := m.titleStyle.Render(fmt.Sprintf("%s: %s", m.layout.Name, m.level.Name))

	return m.contentStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		title,
		m.inputStyle.Render(m.input.View()),
		m.help.View(keys),
	))
}

func (m *Model) SetSize(width int, height int) {
	h, _ := m.contentStyle.GetFrameSize()
	m.input.Width = width - h
}

func (m *Model) SetLayout(layout layout.LayoutCourse) {
	m.layout = layout
}

func (m *Model) SetLevel(level layout.Level) {
	m.level = level
	m.input.Reset()
	m.input.TextToWrite = strings.Join(level.Words, " ")
}
