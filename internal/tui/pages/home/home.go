// Package home renders the Daedalus home page.
package home

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/menu"
	"github.com/vdcds/Daedalus/internal/tui/pages"
)

type Home struct {
	Menu menu.Menu
}

func New() *Home {
	return &Home{
		Menu: menu.Menu{
			Items: []menu.Item{
				{
					ID:          "packages",
					Title:       "Packages",
					Description: "Install CLI tools and GUI applications",
				},
				{
					ID:          "dotfiles",
					Title:       "Dotfiles",
					Description: "Configure your development environment",
				},
				{
					ID:          "fonts",
					Title:       "Fonts",
					Description: "Install Nerd Fonts",
				},
				{
					ID:          "tweaks",
					Title:       "Tweaks",
					Description: "Apply macOS system tweaks",
				},
				{
					ID:          "settings",
					Title:       "Settings",
					Description: "Configure Daedalus",
				},
			},
		},
	}
}

var (
	titleStyle    = lipgloss.NewStyle().Bold(true)
	subtitleStyle = lipgloss.NewStyle().Faint(true)
)

func (h *Home) Update(msg tea.KeyMsg) pages.Page {
	switch msg.String() {

	case "up":
		h.Menu.MoveUp()

	case "down":
		h.Menu.MoveDown()

	case "enter":
		switch h.Menu.SelectedItem().ID {
		case "packages":
			return pages.Packages
		}
	}

	return pages.Home
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
