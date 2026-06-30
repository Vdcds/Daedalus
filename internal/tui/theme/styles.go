package theme

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Title    lipgloss.Style
	Subtitle lipgloss.Style

	Text  lipgloss.Style
	Muted lipgloss.Style

	Selected lipgloss.Style
	Normal   lipgloss.Style

	Border  lipgloss.Style
	Error   lipgloss.Style
	Success lipgloss.Style
}

func NewStyles(p Palette) Styles {
	return Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Rose),

		Subtitle: lipgloss.NewStyle().
			Foreground(p.Muted),

		Text: lipgloss.NewStyle().
			Foreground(p.Text),

		Muted: lipgloss.NewStyle().
			Foreground(p.Muted),

		Selected: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Pine),

		Normal: lipgloss.NewStyle().
			Foreground(p.Text),

		Border: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(p.Overlay),

		Error: lipgloss.NewStyle().
			Foreground(p.Love),

		Success: lipgloss.NewStyle().
			Foreground(p.Foam),
	}
}
