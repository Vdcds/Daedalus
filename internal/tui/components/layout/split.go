// Package layout provides reusable layout primitives.
package layout

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
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

func (s *Split) View() string {
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

	for i := 0; i < height; i++ {
		l := lipgloss.NewStyle().
			Width(s.LeftWidth).
			Render(left[i])

		r := lipgloss.NewStyle().
			Width(s.RightWidth).
			Render(right[i])

		rows = append(rows, l+" │ "+r)
	}

	return strings.Join(rows, "\n")
}
