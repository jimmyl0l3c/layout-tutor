package inputfield

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/runeutil"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	DeleteCharacterBackward key.Binding
}

var DefaultKeyMap = KeyMap{
	DeleteCharacterBackward: key.NewBinding(key.WithKeys("backspace", "ctrl+h")),
}

type Model struct {
	TextToWrite string

	// Styles. These will be applied as inline styles.
	CorrectStyle lipgloss.Style
	ErrorStyle   lipgloss.Style
	TextStyle    lipgloss.Style

	// Width is the maximum number of characters that can be displayed at once.
	// It essentially treats the text field like a horizontally scrolling
	// viewport. If 0 or less this setting is ignored.
	Width int

	// How many characters ahead should be visible during horizontal scrolling.
	ScrollOff int

	KeyMap KeyMap

	// Underlying text value.
	value []rune

	// focus indicates whether user input focus should be on this input
	// component. When false, ignore keyboard input and hide the cursor.
	focus bool

	cursor cursor.Model

	// Cursor position.
	pos int

	// Used to emulate a viewport when width is set and the content is
	// overflowing.
	offset int

	// rune sanitizer for input.
	rsan runeutil.Sanitizer
}

func New() Model {
	return Model{
		CorrectStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#43BF6D")),
		ErrorStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#EB6F92")),
		TextStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("240")),

		KeyMap: DefaultKeyMap,

		ScrollOff: 4,

		cursor: cursor.New(),
		value:  nil,
		focus:  false,
		pos:    0,
	}
}

// Value returns the value of the text input.
func (m Model) Value() string {
	return string(m.value)
}

// Focused returns the focus state on the model.
func (m Model) Focused() bool {
	return m.focus
}

// Focus sets the focus state on the model. When the model is in focus it can
// receive keyboard input and the cursor will be shown.
func (m *Model) Focus() tea.Cmd {
	m.focus = true
	return m.cursor.Focus()
}

// Blur removes the focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (m *Model) Blur() {
	m.focus = false
	m.cursor.Blur()
}

// Reset sets the input to its default state with no input.
func (m *Model) Reset() {
	m.value = nil
	m.pos = 0
	m.offset = 0
}

// rsan initializes or retrieves the rune sanitizer.
func (m *Model) san() runeutil.Sanitizer {
	if m.rsan == nil {
		// Textinput has all its input on a single line so collapse
		// newlines/tabs to single spaces.
		m.rsan = runeutil.NewSanitizer(
			runeutil.ReplaceTabs(" "), runeutil.ReplaceNewlines(" "))
	}
	return m.rsan
}

func (m *Model) handleOverflow() {
	maxLen := len(m.TextToWrite)
	vl := len(m.value)
	scrollOff := min(m.ScrollOff, maxLen-vl)
	available := m.Width - scrollOff

	if maxLen <= m.Width || vl <= available {
		m.offset = 0
		return
	}

	m.offset = vl - available
}

func (m *Model) insertRunesFromUserInput(v []rune) {
	available := len(m.TextToWrite) - len(m.value)

	if available <= 0 {
		return
	}

	// Clean up any special characters in the input provided by the
	// clipboard. This avoids bugs due to e.g. tab characters and
	// whatnot.
	paste := m.san().Sanitize(v)

	limit := min(len(paste), available)
	m.value = append(m.value, paste[:limit]...)

	m.pos = len(m.value)
	m.handleOverflow()
}

// Update is the Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	oldPos := m.pos

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.DeleteCharacterBackward):
			if len(m.value) > 0 {
				m.value = append(m.value[:max(0, m.pos-1)], m.value[m.pos:]...)
				m.pos -= 1
				m.handleOverflow()
			}
		default:
			// Input one or more regular characters.
			m.insertRunesFromUserInput(msg.Runes)
		}

	}

	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.cursor, cmd = m.cursor.Update(msg)
	cmds = append(cmds, cmd)

	if oldPos != m.pos && m.cursor.Mode() == cursor.CursorBlink {
		m.cursor.Blink = false
		cmds = append(cmds, m.cursor.BlinkCmd())
	}

	return m, tea.Batch(cmds...)
}

func (m Model) textToWriteView() string {
	var (
		lv = len(m.value)
		lp = len(m.TextToWrite)
	)

	if lv >= lp {
		return ""
	}

	m.cursor.TextStyle = m.TextStyle
	m.cursor.SetChar(string(m.TextToWrite[lv]))

	textStyle := m.TextStyle.Inline(true).Render

	offset := min(lp, m.Width+m.offset)

	return m.cursor.View() + textStyle(m.TextToWrite[lv+1:offset])
}

func (m Model) validatedTextView() string {
	var (
		correctStyle = m.CorrectStyle.Inline(true).Render
		errorStyle   = m.ErrorStyle.Inline(true).Render
	)

	var sb strings.Builder

	correctText := []rune(m.TextToWrite[m.offset:])

	// Could be improved later to style entire chunks
	for i, r := range m.value[m.offset:] {
		if char := string(r); r == correctText[i] {
			sb.WriteString(correctStyle(char))
		} else {
			sb.WriteString(errorStyle(char))
		}
	}

	return sb.String()
}

// View renders the textinput in its current state.
func (m Model) View() string {
	return m.validatedTextView() + m.textToWriteView()
}

// Blink is a command used to initialize cursor blinking.
func Blink() tea.Msg {
	return cursor.Blink()
}
