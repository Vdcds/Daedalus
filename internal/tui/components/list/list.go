// Package list provides a reusable selectable list component.
package list

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Item struct {
	ID          string
	Title       string
	Description string
	Marked      bool
}

type List struct {
	Items    []Item
	Selected int
}

func (l *List) Update(msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		if l.Selected > 0 {
			l.Selected--
		}

	case "down":
		if l.Selected < len(l.Items)-1 {
			l.Selected++
		}
	}
}

func (l List) SelectedItem() Item {
	if len(l.Items) == 0 {
		return Item{}
	}

	return l.Items[l.Selected]
}

func (l List) View(t theme.Theme) string {
	var out []string

	for i, item := range l.Items {
		selected := i == l.Selected

		cursor := "  "
		marker := "  "
		titleStyle := t.Styles.Normal

		if selected {
			cursor = t.Styles.Highlight.Render("▎ ")
			titleStyle = t.Styles.Title
		}

		if item.Marked {
			marker = t.Styles.Highlight.Render("✓ ")
		}

		out = append(
			out,
			cursor+marker+titleStyle.Render(item.Title),
		)

		if item.Description != "" {
			out = append(
				out,
				"     "+t.Styles.Muted.Render(item.Description),
			)
		}

		if i != len(l.Items)-1 {
			out = append(out, "")
		}
	}

	return strings.Join(out, "\n")
}
