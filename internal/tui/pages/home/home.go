// Package home renders the Daedalus home page.
package home

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	"github.com/vdcds/Daedalus/internal/tui/components/menu"
	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Home struct {
	Header *header.Header
	Menu   menu.Menu
	Footer *footer.Footer
}

func New() *Home {
	return &Home{
		Header: header.New(
			"🪽 Daedalus",
			"Forge your macOS experience",
		),

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

		Footer: footer.New(
			footer.Action{
				Key:         "↑↓",
				Description: "Navigate",
			},
			footer.Action{
				Key:         "Enter",
				Description: "Select",
			},
			footer.Action{
				Key:         "Q",
				Description: "Quit",
			},
		),
	}
}

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

func (h *Home) View(t theme.Theme) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,

		h.Header.View(t),

		"",

		h.Menu.View(t),

		"",

		h.Footer.View(t),
	)
}
