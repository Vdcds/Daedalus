// Package menu provides a reusable selectable menu component.
package menu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	ID          string
	Title       string
	Description string
}

type Menu struct {
	Items    []Item
	Selected int
}

var (
	selectedStyle = lipgloss.NewStyle().Bold(true)
	normalStyle   = lipgloss.NewStyle()
)

func (m *Menu) MoveUp() {
	if m.Selected > 0 {
		m.Selected--
	}
}

func (m *Menu) MoveDown() {
	if m.Selected < len(m.Items)-1 {
		m.Selected++
	}
}

func (m Menu) SelectedItem() Item {
	return m.Items[m.Selected]
}

func (m Menu) View() string {
	var lines []string

	for i, item := range m.Items {
		if i == m.Selected {
			lines = append(lines, selectedStyle.Render("❯ "+item.Title))
		} else {
			lines = append(lines, normalStyle.Render("  "+item.Title))
		}
	}

	return strings.Join(lines, "\n")
}
