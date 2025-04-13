package course

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Back  key.Binding
	Reset key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Reset, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Back, k.Reset, k.Quit}}
}

var keys = keyMap{
	Back: key.NewBinding(
		key.WithKeys("-"),
		key.WithHelp("-", "back to menu"),
	),
	Reset: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "reset level"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
