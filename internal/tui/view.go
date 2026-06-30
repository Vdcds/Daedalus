package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Faint(true)
)

func (a *App) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("🪽 Daedalus"),

		"",

		subtitleStyle.Render("Forge your macOS experience"),

		"",

		subtitleStyle.Render("(press q to quit)"),
	)

	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
