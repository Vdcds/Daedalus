// Package header provides a compact breadcrumb-style page header.
package header

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Header renders a single-line title · subtitle breadcrumb
// with optional right-aligned content.
type Header struct {
	Title    string
	Subtitle string
	Right    string
}

func New(title, subtitle string) *Header {
	return &Header{Title: title, Subtitle: subtitle}
}

func (h *Header) SetRight(content string) {
	h.Right = content
}

func (h *Header) View(t theme.Theme) string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(t.Palette.Iris).
		Render(h.Title)

	line := title
	if h.Subtitle != "" {
		sep := lipgloss.NewStyle().Foreground(t.Palette.Overlay).Render(" · ")
		sub := lipgloss.NewStyle().Foreground(t.Palette.Muted).Render(h.Subtitle)
		line = title + sep + sub
	}

	if h.Right == "" {
		return line
	}

	leftW := 56
	rightW := 30

	left := lipgloss.NewStyle().Width(leftW).Render(line)
	right := lipgloss.NewStyle().Width(rightW).Align(lipgloss.Right).Render(h.Right)

	return lipgloss.JoinHorizontal(lipgloss.Center, left, right)
}
