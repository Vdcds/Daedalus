// Package home renders the Daedalus home page.
package home

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/menu"

	tea "github.com/charmbracelet/bubbletea"
)

type Home struct {
	Menu menu.Menu
}

func New() *Home {
	return &Home{
		Menu: menu.Menu{
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
		},
	}
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Faint(true)
)

func (h *Home) Update(msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		h.Menu.MoveUp()

	case "down":
		h.Menu.MoveDown()
	}
}

func (h *Home) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("🪽 Daedalus"),

		"",

		subtitleStyle.Render("Forge your macOS experience"),

		"",

		h.Menu.View(),
	)
}
