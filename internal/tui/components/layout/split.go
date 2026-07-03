// Package layout provides reusable layout primitives.
package layout

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Split struct {
	Left  string
	Right string

	LeftWidth  int
	RightWidth int
}

func NewSplit(left, right string) *Split {
	return &Split{
		Left:       left,
		Right:      right,
		LeftWidth:  30,
		RightWidth: 61,
	}
}

func (s *Split) View(t theme.Theme) string {
	left := strings.Split(s.Left, "\n")
	right := strings.Split(s.Right, "\n")

	height := len(left)
	if len(right) > height {
		height = len(right)
	}

	for len(left) < height {
		left = append(left, "")
	}

	for len(right) < height {
		right = append(right, "")
	}

	var rows []string

	divider := t.Styles.Muted.Render("│")

	for i := 0; i < height; i++ {
		l := lipgloss.NewStyle().
			Width(s.LeftWidth).
			Render(left[i])

		r := lipgloss.NewStyle().
			Width(s.RightWidth).
			Render(right[i])

		rows = append(rows, l+" "+divider+" "+r)
	}

	return strings.Join(rows, "\n")
}
