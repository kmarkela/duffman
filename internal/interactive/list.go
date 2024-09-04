package interactive

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kmarkela/duffman/internal/pcollection"
)

const listHeight = 14

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

	if i.Node != nil {
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
	stack    []item   // Stack to keep track of node levels
	tstack   []item   // temp Stack to keep track of node levels
	path     []string // To keep the current path for display
	back     bool     // going backwards
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
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
			if ok && len(i.Node) > 0 { // If selected item has a sublist

				m.tstack = m.stack
				m.stack = append(m.stack, i)    // Push current items to stack
				m.path = append(m.path, i.Name) // Update path
				m.updateList(i)
			} else if ok {
				fmt.Println("Selected sublist item:", i.Name)
				return m, tea.Quit
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
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	// var ll []string
	// for _, k := range m.stack {
	// 	ll = append(ll, k.Name)
	// }

	// header := fmt.Sprintf("\nCurrent Path: %s, Stack: %s\n", strings.Join(m.path, " > "), strings.Join(ll, " > ")) // Display current path
	header := fmt.Sprintf("\nCurrent Path: %s\n", strings.Join(m.path, " > ")) // Display current path
	return header + "\n" + m.list.View()
}

// Function to update the list model with new items
func (m *model) updateList(i item) {
	items := []list.Item{}

	for _, k := range i.Node {
		items = append(items, item(k))
	}

	m.list.SetItems(items)
	m.list.CursorUp()
}

func RenderList(nl pcollection.NodeList) {
	items := []list.Item{}
	fmt.Println(len(nl))
	for _, k := range nl {
		items = append(items, item(k))
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want for dinner?"
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l, stack: make([]item, 0), path: []string{"Root"}}

	if _, err := tea.NewProgram(&m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
