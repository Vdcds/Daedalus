// Package card provides a reusable bordered panel component.
package card

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Card is a bordered panel with an internal title line.
type Card struct {
	Title   string
	Body    string
	Width   int
	Height  int
	Focused bool
}

// New creates a card with the given title, pre-rendered body, and width.
// Width sets the content+padding width (borders are extra).
func New(title, body string, width int) *Card {
	return &Card{Title: title, Body: body, Width: width}
}

// WithHeight sets the content+padding height (borders are extra).
func (c *Card) WithHeight(height int) *Card {
	c.Height = height
	return c
}

// WithFocus marks the panel as the active one.
func (c *Card) WithFocus(focused bool) *Card {
	c.Focused = focused
	return c
}

func (c *Card) View(t theme.Theme) string {
	borderFg := t.Palette.Overlay
	titleFg := t.Palette.Text
	if c.Focused {
		borderFg = t.Palette.Pine
		titleFg = t.Palette.Iris
	}

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(titleFg).
		Render(c.Title)

	// Title on first line, blank separator, then body content.
	content := title + "\n\n" + c.Body

	style := lipgloss.NewStyle().
		Width(c.Width).
		Padding(1, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderFg)

	if c.Height > 0 {
		style = style.Height(c.Height)
	}

	return style.Render(content)
}
