package interactive

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/kmarkela/duffman/internal/logger"
)

const (
	initialInputs = 3
	helpHeight    = 5
)

var (
	editorTitleStyle = lipgloss.NewStyle().MarginLeft(0).Bold(true)

	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	cursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("57")).
			Foreground(lipgloss.Color("230"))

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	endOfBufferStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("235"))

	focusedPlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99"))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.HiddenBorder())
)

type keymap struct {
	next, send, back, quit key.Binding
}

func newTextarea() textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return t
}

type modelEditor struct {
	width  int
	height int
	keymap keymap
	help   help.Model
	inputs []textarea.Model
	focus  int
	ml     *model
}

func newModel(i item, ml *model) modelEditor {
	me := modelEditor{
		inputs: make([]textarea.Model, initialInputs),
		help:   help.New(),
		ml:     ml,
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "next"),
			),
			send: key.NewBinding(
				key.WithKeys("ctrl+r"),
				key.WithHelp("ctrl+r", "send Req"),
			),
			back: key.NewBinding(
				key.WithKeys("ctrl+l"),
				key.WithHelp("ctrl+l", "back"),
			),
			quit: key.NewBinding(
				key.WithKeys("esc", "ctrl+c"),
				key.WithHelp("esc", "quit"),
			),
		},
	}
	for i := 0; i < initialInputs; i++ {
		me.inputs[i] = newTextarea()
	}
	me.inputs[me.focus].Focus()

	me.inputs[0].SetValue(buildReqStr(*i.Req))

	v := buildVarStr(*ml.col)

	logger.Logger.Info(v)

	me.inputs[1].SetValue(buildVarStr(*ml.col))

	width, height, _ := term.GetSize(0)
	me.height = height
	me.width = width
	me.sizeInputs()

	return me
}

func (m modelEditor) Init() tea.Cmd {
	return textarea.Blink

}

func (m modelEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height, _ := term.GetSize(0)
		m.height = height
		m.width = width
		m.sizeInputs()

	case tea.KeyMsg:

		switch {
		case key.Matches(msg, m.keymap.quit):
			for i := range m.inputs {
				m.inputs[i].Blur()
			}
			return m, tea.Quit
		case key.Matches(msg, m.keymap.back):
			return m.ml, nil
		case key.Matches(msg, m.keymap.next):
			m.inputs[m.focus].Blur()
			m.focus++
			if m.focus > len(m.inputs)-1 {
				m.focus = 0
			}
			cmd := m.inputs[m.focus].Focus()
			cmds = append(cmds, cmd)
		}
	}

	// Update all textareas
	for i := range m.inputs {
		newModel, cmd := m.inputs[i].Update(msg)
		m.inputs[i] = newModel
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *modelEditor) sizeInputs() {

	for i := range m.inputs {
		m.inputs[i].SetWidth(m.width / 2)
	}

	h := (m.height - helpHeight) / 4
	m.inputs[0].SetHeight(h*3 - 2)
	m.inputs[1].SetHeight(h)

	m.inputs[2].SetHeight(m.height - helpHeight - 1)
}

func (m modelEditor) View() string {

	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		// m.keymap.prev,
		m.keymap.send,
		m.keymap.back,
		m.keymap.quit,
	})

	var views []string

	// Combine title and input view
	editorView := lipgloss.JoinVertical(lipgloss.Top,
		editorTitleStyle.Render("REQUEST:"),
		m.inputs[0].View(),
		editorTitleStyle.Render("VARIABLES:"),
		m.inputs[1].View())

	views = append(views, editorView)

	// Combine title and input view
	editorView = lipgloss.JoinVertical(lipgloss.Top,
		editorTitleStyle.Render("RESPONSE:"),
		m.inputs[2].View())
	views = append(views, editorView)

	return lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n" + help
}
