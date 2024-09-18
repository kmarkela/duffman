package interactive

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/internal/req"
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
	next, send, back, save, quit key.Binding
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
	item   item
}

func newModel(i item, ml *model) modelEditor {
	me := modelEditor{
		inputs: make([]textarea.Model, initialInputs),
		help:   help.New(),
		ml:     ml,
		item:   i,
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
			save: key.NewBinding(
				key.WithKeys("ctrl+s"),
				key.WithHelp("ctrl+s", "save vars"),
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

	width, height, _ := term.GetSize(0)
	me.height = height
	me.width = width
	me.sizeInputs()

	req := buildReqStr(*i.Req, ml.col.Env, ml.col.Variables)
	me.inputs[0].CharLimit = 2 * len(req)
	me.inputs[0].MaxHeight = len(strings.Split(req, "\n"))
	me.inputs[0].SetValue(req)

	vars := buildVarStr(*ml.col)
	me.inputs[1].CharLimit = 2 * len(vars)
	me.inputs[1].MaxHeight = len(strings.Split(vars, "\n"))
	me.inputs[1].SetValue(vars)

	return me
}

func (m modelEditor) Init() tea.Cmd {
	return textarea.Blink

}

func (m modelEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// TODO: refactor
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
		case key.Matches(msg, m.keymap.save):
			m.saveVars()
		case key.Matches(msg, m.keymap.next):
			m.inputs[m.focus].Blur()
			m.focus++
			if m.focus > len(m.inputs)-1 {
				m.focus = 0
			}
			cmd := m.inputs[m.focus].Focus()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keymap.send):
			m.sendReq()
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

func (m *modelEditor) saveVars() {
	var vo varOut

	if err := json.Unmarshal([]byte(m.inputs[1].Value()), &vo); err != nil {
		m.inputs[0].SetValue(fmt.Sprintf("\nError of parsing VARIABLES. Error Msg: %s", err.Error()))
		return
	}

	req := buildReqStr(*m.item.Req, vo.Env, vo.Variables)
	m.inputs[0].SetValue(req)

	// save vars
	m.ml.col.Env = vo.Env
	m.ml.col.Variables = vo.Variables
}

func (m *modelEditor) sendReq() {
	var ro pcollection.Req

	if err := json.Unmarshal([]byte(m.inputs[0].Value()), &ro); err != nil {
		m.inputs[0].SetValue(fmt.Sprintf("\nError of parsing REQUEST. Error Msg: %s", err.Error()))
		return
	}

	var body io.Reader = strings.NewReader(ro.Body)
	res, err := req.DoRequestFull(ro.URL, body, ro, m.ml.tr)
	if err != nil {
		m.inputs[2].SetValue(fmt.Sprintf("\nError executing request. Error Msg: %s", err.Error()))
		return
	}
	var header string = res.Status
	for k, v := range res.Header {
		header = fmt.Sprintf("%s\n%s: %s", header, k, strings.Join(v, ";"))
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		m.inputs[2].SetValue(fmt.Sprintf("\nError executing request. Error Msg: %s", err.Error()))
	}
	res.Body.Close()

	m.inputs[2].CharLimit = len(string(bodyBytes)) + len(header) + 2
	m.inputs[2].MaxHeight = len(strings.Split(string(bodyBytes), "\n")) + len(res.Header) + 1

	m.inputs[2].SetValue(fmt.Sprintf("%s\n\n%s", header, string(bodyBytes)))
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
		m.keymap.save,
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
