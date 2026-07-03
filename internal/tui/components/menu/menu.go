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
	var out []string

	for i, item := range m.Items {
		selected := i == m.Selected

		cursor := "  "
		titleStyle := t.Styles.Normal

		if selected {
			cursor = t.Styles.Highlight.Render("▎ ")
			titleStyle = t.Styles.Title
		}

		out = append(out,
			cursor+titleStyle.Render(item.Title),
		)

		if item.Description != "" {
			out = append(out,
				"   "+t.Styles.Muted.Render(item.Description),
			)
		}

		if i != len(m.Items)-1 {
			out = append(out, "")
		}
	}

	return strings.Join(out, "\n")
}
