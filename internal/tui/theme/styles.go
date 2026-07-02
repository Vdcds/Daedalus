package theme

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Title    lipgloss.Style
	Subtitle lipgloss.Style

	Text      lipgloss.Style
	Muted     lipgloss.Style
	Highlight lipgloss.Style

	Selected lipgloss.Style
	Normal   lipgloss.Style

	Accent lipgloss.Style

	Help lipgloss.Style

	Border lipgloss.Style

	Error   lipgloss.Style
	Success lipgloss.Style
}

func NewStyles(p Palette) Styles {
	return Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Iris),

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

		Accent: lipgloss.NewStyle().
			Foreground(p.Rose),

		Help: lipgloss.NewStyle().
			Faint(true).
			Foreground(p.Muted),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(p.Overlay).
			Padding(0, 1),

		Error: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Love),

		Success: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Foam),
		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Highlight),
	}
}
