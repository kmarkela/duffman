package interactive

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

// Simple tab view model for items without sub-nodes
type tabView struct {
	Tabs      []string // Tab names
	activeTab int      // Currently active tab index
	content   string   // Content to display
	width     int
	height    int
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func newTabView(itemName string) *tabView {
	return &tabView{
		Tabs:      []string{"Info", "Details", "Stats"}, // Example tabs
		activeTab: 0,
		content:   fmt.Sprintf("Content for: %s", itemName),
	}
}

func (t tabView) Init() tea.Cmd {
	return nil
}

func (t *tabView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.list.SetWidth(msg.Width)
	// 	m.list.SetHeight(msg.Height)
	// 	return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return t, tea.Quit
		case "tab", "right":
			t.activeTab = min(t.activeTab+1, len(t.Tabs)-1) // Move to next tab
			// return t, nil
		case "shift+tab", "left":
			t.activeTab = max(t.activeTab-1, 0) // Move to previous tab
			// return t, nil
		case "esc", "backspace":
			return t, tea.Quit // Return to the previous view
		}
	}
	return t, nil
}

func (t tabView) View() string {
	// Render the tabs
	doc := strings.Builder{}

	var renderedTabs []string

	for i, k := range t.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(t.Tabs)-1, i == t.activeTab
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(k))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(t.content))
	return docStyle.Render(doc.String())
}
