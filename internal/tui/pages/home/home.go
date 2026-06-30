// Package home renders the Daedalus Home Screeen.
package home

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Faint(true)
)

func View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("🪽 Daedalus"),

		"",

		subtitleStyle.Render("Forge your macOS experience"),

		"",

		subtitleStyle.Render("(press q to quit)"),
	)
}
