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

var MainMenu = menu.Menu{
	Items: []menu.Item{
		{
			Title:       "Packages",
			Description: "Install CLI tools and GUI applications",
		},
		{
			Title:       "Dotfiles",
			Description: "Configure your development environment",
		},
		{
			Title:       "Fonts",
			Description: "Install Nerd Fonts",
		},
		{
			Title:       "Tweaks",
			Description: "Apply macOS system tweaks",
		},
		{
			Title:       "Settings",
			Description: "Configure Daedalus",
		},
	},
}

func View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("🪽 Daedalus"),

		"",

		subtitleStyle.Render("Forge your macOS experience"),

		"",

		MainMenu.View(),
	)
}
