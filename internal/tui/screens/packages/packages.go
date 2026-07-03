// Package packages renders the packages screen.
package packages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	"github.com/vdcds/Daedalus/internal/tui/components/menu"
	"github.com/vdcds/Daedalus/internal/tui/components/preview"
	"github.com/vdcds/Daedalus/internal/tui/components/search"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Packages struct {
	Header  *header.Header
	Search  *search.Search
	Menu    menu.Menu
	Preview *preview.Preview
	Footer  *footer.Footer
}

func New() *Packages {
	s := search.New("Search packages...")
	s.Focus()

	return &Packages{
		Header: header.New(
			"📦 Packages",
			"Browse package categories",
		),

		Search: s,

		Menu: menu.Menu{
			Items: []menu.Item{
				{
					ID:          "cli",
					Title:       "CLI Tools",
					Description: "Core terminal utilities",
				},
				{
					ID:          "editors",
					Title:       "Editors",
					Description: "Code editors & IDEs",
				},
				{
					ID:          "browsers",
					Title:       "Browsers",
					Description: "Firefox, Zen, Chromium...",
				},
				{
					ID:          "development",
					Title:       "Development",
					Description: "Git, Docker, Go, Node.js",
				},
				{
					ID:          "utilities",
					Title:       "Utilities",
					Description: "Everyday applications",
				},
			},
		},

		Preview: preview.New(
			"Welcome",
			"Select a category to begin browsing.",
			"",
			"Package information will appear here.",
		),

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
	p.Search.Update(msg)
	p.Menu.Update(msg)

	switch msg.String() {

	case "esc":
		return screens.Home

	case "enter":
		// Category navigation later.
	}

	return screens.Packages
}

func (p *Packages) View(t theme.Theme) string {
	p.Header.SetRight(p.Search.View(t))

	left := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render("Categories"),
		"",
		p.Menu.View(t),
	)

	right := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render("Preview"),
		"",
		p.Preview.View(t),
	)

	body := layout.NewSplit(
		left,
		right,
	).View()

	return layout.New(
		p.Header.View(t),
		body,
		p.Footer.View(t),
	).View()
}
