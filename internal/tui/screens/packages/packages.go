// Package packages renders the packages screen.
package packages

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	"github.com/vdcds/Daedalus/internal/tui/components/menu"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Packages struct {
	Header *header.Header
	Menu   menu.Menu
	Footer *footer.Footer
}

func New() *Packages {
	return &Packages{
		Header: header.New(
			"📦 Packages",
			"Browse package categories",
		),

		Menu: menu.Menu{
			Items: []menu.Item{
				{
					ID:          "cli",
					Title:       "CLI Tools",
					Description: "Essential command-line utilities",
				},
				{
					ID:          "editors",
					Title:       "Editors",
					Description: "Neovim, Zed, VS Code and more",
				},
				{
					ID:          "browsers",
					Title:       "Browsers",
					Description: "Firefox, Zen, Brave, Chromium",
				},
				{
					ID:          "development",
					Title:       "Development",
					Description: "Git, Docker, Node.js, Go",
				},
				{
					ID:          "utilities",
					Title:       "Utilities",
					Description: "Everyday desktop applications",
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
				Description: "Open",
			},
			footer.Action{
				Key:         "Esc",
				Description: "Back",
			},
		),
	}
}

func (p *Packages) Update(msg tea.KeyMsg) screens.Screen {
	p.Menu.Update(msg)

	switch msg.String() {

	case "esc":
		return screens.Home

	case "enter":
		// Category navigation will be implemented later.
	}

	return screens.Packages
}

func (p *Packages) View(t theme.Theme) string {
	return layout.New(
		p.Header.View(t),
		p.Menu.View(t),
		p.Footer.View(t),
	).View()
}
