// Package layout provides application-wide layout primitives.
package layout

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	MaxWidth          = 96
	HorizontalPadding = 2
	VerticalPadding   = 1

	SectionSpacing = 1
)

type Layout struct {
	Header string
	Body   string
	Footer string
}

func New(header, body, footer string) *Layout {
	return &Layout{
		Header: header,
		Body:   body,
		Footer: footer,
	}
}

func (l *Layout) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,

		l.Header,

		strings.Repeat("\n", SectionSpacing),

		l.Body,

		strings.Repeat("\n", SectionSpacing),

		l.Footer,
	)

	return lipgloss.NewStyle().
		Width(MaxWidth).
		Padding(
			VerticalPadding,
			HorizontalPadding,
		).
		Render(content)
}
