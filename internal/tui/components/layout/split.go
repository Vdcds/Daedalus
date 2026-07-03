// Package layout provides reusable layout primitives.
package layout

import "github.com/charmbracelet/lipgloss"

type Split struct {
	Left  string
	Right string

	LeftWidth  int
	RightWidth int

	Gap int
}

func NewSplit(left, right string) *Split {
	return &Split{
		Left:       left,
		Right:      right,
		LeftWidth:  30,
		RightWidth: 62,
		Gap:        4,
	}
}

func (s *Split) View() string {
	left := lipgloss.NewStyle().
		Width(s.LeftWidth).
		Render(s.Left)

	right := lipgloss.NewStyle().
		Width(s.RightWidth).
		Render(s.Right)

	gap := lipgloss.NewStyle().
		Width(s.Gap).
		Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		gap,
		right,
	)
}
