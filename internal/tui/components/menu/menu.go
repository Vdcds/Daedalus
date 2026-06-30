// Package menu provides a reusable selectable menu component.
package menu

import (
	"strings"

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

func (m Menu) View(t theme.Theme) string {
	var lines []string

	for i, item := range m.Items {

		if i == m.Selected {
			lines = append(lines,
				t.Styles.Selected.Render("❯ "+item.Title),
			)
		} else {
			lines = append(lines,
				t.Styles.Normal.Render("  "+item.Title),
			)
		}

	}

	return strings.Join(lines, "\n")
}
