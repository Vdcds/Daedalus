// Package menu provides a reusable selectable menu component.
package menu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	Title string
}

type Menu struct {
	Items    []Item
	Selected int
}

var (
	selectedStyle = lipgloss.NewStyle().
			Bold(true)

	normalStyle = lipgloss.NewStyle()
)

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
