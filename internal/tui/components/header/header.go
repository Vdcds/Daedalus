// Package header provides a reusable page header.
package header

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Header struct {
	Title    string
	Subtitle string

	Right string
}

func New(title, subtitle string) *Header {
	return &Header{
		Title:    title,
		Subtitle: subtitle,
	}
}

func (h *Header) SetRight(right string) {
	h.Right = right
}

func (h *Header) View(t theme.Theme) string {
	left := lipgloss.JoinVertical(
		lipgloss.Left,

		t.Styles.Title.Render(h.Title),

		t.Styles.Subtitle.Render(h.Subtitle),
	)

	right := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Width(32).
		Render(h.Right)

	top := lipgloss.JoinHorizontal(
		lipgloss.Top,

		lipgloss.NewStyle().
			Width(58).
			Render(left),

		right,
	)

	divider := t.Styles.Muted.Render(
		"────────────────────────────────────────────────────────────────────────────────────────────",
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,

		top,

		divider,
	)
}
