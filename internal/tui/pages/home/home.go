// Package home renders the Daedalus home page.
package home

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/menu"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Faint(true)
)

func View() string {
	mainMenu := menu.Menu{
		Items: []menu.Item{
			{Title: "Packages"},
			{Title: "Dotfiles"},
			{Title: "Fonts"},
			{Title: "Tweaks"},
			{Title: "Settings"},
		},
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("🪽 Daedalus"),

		"",

		subtitleStyle.Render("Forge your macOS experience"),

		"",

		mainMenu.View(),
	)
}
