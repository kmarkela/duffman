package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kmarkela/duffman/internal/pcollection"
)

// const listHeight = 140

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(0)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item pcollection.Node

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	if i.Req == nil {
		i.Name = fmt.Sprintf("📁 %s", i.Name) // Display folders with an icon
	}

	str := fmt.Sprintf("[ ] %s", i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(strings.ReplaceAll(strings.Join(s, " "), "[ ]", "[x]"))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	quitting bool
	stack    []item                  // Stack to keep track of node levels
	tstack   []item                  // temp Stack to keep track of node levels
	path     []string                // To keep the current path for display
	back     bool                    // going backwards
	col      *pcollection.Collection // TODO: DO I need whole collection here?
	tr       *http.Transport
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if len(m.stack) == 0 {

				ti := item{}
				for _, k := range m.list.Items() {
					ti.Node = append(ti.Node, pcollection.Node{Name: k.(item).Name, Node: k.(item).Node, Req: k.(item).Req})
				}
				m.stack = append(m.stack, ti)
			}

			i, ok := m.list.SelectedItem().(item)

			if m.back && len(m.stack) < len(m.path) {
				m.stack = append(m.stack, m.tstack[len(m.tstack)-2])
			}
			m.back = false

			if ok && i.Req == nil { // If selected item has a sublist

				m.tstack = m.stack
				m.stack = append(m.stack, i)    // Push current items to stack
				m.path = append(m.path, i.Name) // Update path
				m.updateList(i)
			} else if ok {
				return newModel(i, m), nil
			}

		case "backspace", "esc":
			if len(m.stack) > 0 {

				m.tstack = m.stack

				// Remove self from stack
				if !m.back {
					m.stack = m.stack[:len(m.stack)-1]
				}
				m.back = true

				m.path = m.path[:len(m.path)-1] // Update path
				last := m.stack[len(m.stack)-1] // Get last items from stack
				m.stack = m.stack[:len(m.stack)-1]
				m.updateList(last)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Exiting...")
	}

	header := fmt.Sprintf("\nCurrent Path: %s\n", strings.Join(m.path, " > ")) // Display current path
	return header + "\n" + m.list.View()
}

// Function to update the list model with new items
func (m *model) updateList(i item) {
	items := []list.Item{}

	for _, k := range i.Node {
		items = append(items, item(k))
	}

	// CursorUp to the first position
	for m.list.Cursor() > 0 {
		m.list.CursorUp()
	}

	m.list.SetItems(items)

}

func additionalShortHelpKeys() []key.Binding {
	var kbl []key.Binding

	kbl = append(kbl, key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "back"),
	))
	return kbl
}

func (c *Client) RenderList(col *pcollection.Collection) {
	items := []list.Item{}
	for _, k := range col.Schema.Nodes {
		items = append(items, item(k))
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, 30)

	l.Title = col.Schema.Name
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.InfiniteScrolling = true
	l.AdditionalShortHelpKeys = additionalShortHelpKeys

	m := model{
		list:  l,
		col:   col,
		stack: make([]item, 0),
		path:  []string{"Root"},
		tr:    c.tr}

	if _, err := tea.NewProgram(&m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
