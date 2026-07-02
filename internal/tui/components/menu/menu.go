// Package menu provides a reusable selectable menu component.
package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/theme"
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

func (m *Menu) Update(msg tea.KeyMsg) {
	switch msg.String() {

	case "up":
		if m.Selected > 0 {
			m.Selected--
		}

	case "down":
		if m.Selected < len(m.Items)-1 {
			m.Selected++
		}
	}
}

func (m Menu) SelectedItem() Item {
	return m.Items[m.Selected]
}

func (m Menu) View(t theme.Theme) string {
	var lines []string

	for i, item := range m.Items {
		lines = append(lines, m.renderItem(item, i == m.Selected, t))
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func (m Menu) renderItem(item Item, selected bool, t theme.Theme) string {
	var lines []string

	title := "  " + item.Title

	if selected {
		title = "❯ " + item.Title
		lines = append(lines, t.Styles.Selected.Render(title))
	} else {
		lines = append(lines, t.Styles.Normal.Render(title))
	}

	if item.Description != "" {
		lines = append(
			lines,
			t.Styles.Muted.Render("    "+item.Description),
		)
	}

	return strings.Join(lines, "\n")
}
