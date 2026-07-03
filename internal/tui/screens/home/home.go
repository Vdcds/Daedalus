// Package home renders the Daedalus home screen.
package home

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	list "github.com/vdcds/Daedalus/internal/tui/components/list"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Home struct {
	Header *header.Header
	List   list.List
	Footer *footer.Footer
}

func New() *Home {
	return &Home{
		Header: header.New(
			"🪽 Daedalus",
			"Forge your macOS experience",
		),

		List: list.List{
			Items: []list.Item{
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

func (h *Home) Update(msg tea.KeyMsg) screens.Screen {
	h.List.Update(msg)

	switch msg.String() {

	case "enter":
		switch h.List.SelectedItem().ID {

		case "packages":
			return screens.Packages
		}
	}

	return screens.Home
}

func (h *Home) View(t theme.Theme) string {
	return layout.New(
		h.Header.View(t),
		h.List.View(t),
		h.Footer.View(t),
	).View()
}
