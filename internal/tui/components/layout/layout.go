// Package layout provides the application layout frame.
package layout

import "github.com/charmbracelet/lipgloss"

// MaxWidth is the outer width of the layout content area (excluding margins).
const MaxWidth = 92

// Layout is a vertical stack of header, body, and footer.
type Layout struct {
	Header string
	Body   string
	Footer string
}

func New(header, body, footer string) *Layout {
	return &Layout{Header: header, Body: body, Footer: footer}
}

func (l *Layout) View() string {
	header := lipgloss.NewStyle().
		MarginBottom(1).
		Render(l.Header)

	body := lipgloss.NewStyle().
		MarginBottom(1).
		Render(l.Body)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		l.Footer,
	)

	return lipgloss.NewStyle().
		Width(MaxWidth).
		Padding(1, 2).
		Render(content)
}
