package menu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list  list.Model
	style lipgloss.Style
}

func New(title string, items []list.Item) Model {
	m := Model{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),

		style: lipgloss.NewStyle().Margin(1, 2),
	}
	m.list.Title = title
	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.style.Render(m.list.View())
}

func (m Model) GetSelected() list.Item {
	m.list.GlobalIndex()
	return m.list.SelectedItem()
}

func (m Model) GetSelectedIndex() int {
	return m.list.GlobalIndex()
}

func (m *Model) SetItems(items []list.Item) tea.Cmd {
	return m.list.SetItems(items)
}

func (m *Model) SetSize(width int, height int) {
	h, v := m.style.GetFrameSize()
	m.list.SetSize(width-h, height-v)
}
