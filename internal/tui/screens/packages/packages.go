// Package packages renders the packages screen.
package packages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/components/card"
	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	"github.com/vdcds/Daedalus/internal/tui/components/menu"
	"github.com/vdcds/Daedalus/internal/tui/components/preview"
	"github.com/vdcds/Daedalus/internal/tui/components/search"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// categoryInfo holds preview data for each category.
type categoryInfo struct {
	Description string
	Packages    []string
	Sections    []preview.Section
}

var categories = map[string]categoryInfo{
	"cli": {
		Description: "Essential command-line utilities for productivity and system management.",
		Packages:    []string{"bat", "eza", "fd", "fzf", "ripgrep", "jq", "htop", "tldr"},
		Sections: []preview.Section{
			{Label: "Count", Value: "24 formulae"},
			{Label: "Source", Value: "Homebrew"},
			{Label: "Install", Value: "brew install <pkg>"},
		},
	},
	"editors": {
		Description: "Code editors and IDEs for every workflow and language.",
		Packages:    []string{"neovim", "zed", "visual-studio-code", "sublime-text"},
		Sections: []preview.Section{
			{Label: "Count", Value: "6 casks"},
			{Label: "Source", Value: "Homebrew Cask"},
			{Label: "Install", Value: "brew install --cask <pkg>"},
		},
	},
	"browsers": {
		Description: "Modern web browsers for daily browsing and development.",
		Packages:    []string{"firefox", "zen-browser", "brave-browser", "chromium"},
		Sections: []preview.Section{
			{Label: "Count", Value: "4 casks"},
			{Label: "Source", Value: "Homebrew Cask"},
			{Label: "Install", Value: "brew install --cask <pkg>"},
		},
	},
	"development": {
		Description: "Developer tools, runtimes, and version managers.",
		Packages:    []string{"git", "docker", "node", "go", "rustup", "python"},
		Sections: []preview.Section{
			{Label: "Count", Value: "12 formulae"},
			{Label: "Source", Value: "Homebrew"},
			{Label: "Install", Value: "brew install <pkg>"},
		},
	},
	"utilities": {
		Description: "Everyday desktop applications for file management, media, and more.",
		Packages:    []string{"raycast", "rectangle", "alt-tab", "stats", "iina"},
		Sections: []preview.Section{
			{Label: "Count", Value: "8 casks"},
			{Label: "Source", Value: "Homebrew Cask"},
			{Label: "Install", Value: "brew install --cask <pkg>"},
		},
	},
}

type Packages struct {
	Header  *header.Header
	Search  *search.Search
	Menu    menu.Menu
	Preview *preview.Preview
	Footer  *footer.Footer
}

func New() *Packages {
	s := search.New("filter…")
	s.Focus()

	p := &Packages{
		Header: header.New(
			"📦 Packages",
			"Browse categories",
		),

		Search: s,

		Menu: menu.Menu{
			Items: []menu.Item{
				{ID: "cli", Title: "CLI Tools", Description: "bat, eza, fd, fzf, rg…", Badge: "24"},
				{ID: "editors", Title: "Editors", Description: "nvim, zed, code…", Badge: "6"},
				{ID: "browsers", Title: "Browsers", Description: "firefox, zen, brave…", Badge: "4"},
				{ID: "development", Title: "Development", Description: "git, docker, node, go…", Badge: "12"},
				{ID: "utilities", Title: "Utilities", Description: "raycast, rectangle…", Badge: "8"},
			},
		},

		Preview: preview.New(""),
		Footer: footer.New(
			footer.Action{Key: "j/k", Description: "navigate"},
			footer.Action{Key: "enter", Description: "open"},
			footer.Action{Key: "/", Description: "filter"},
			footer.Action{Key: "esc", Description: "back"},
		),
	}

	// Set initial preview from first category.
	p.syncPreview()
	return p
}

func (p *Packages) syncPreview() {
	sel := p.Menu.SelectedItem()
	info, ok := categories[sel.ID]
	if !ok {
		return
	}

	p.Preview.SetTitle(sel.Title)
	p.Preview.SetContent(info.Description)

	// Build sections: metadata + package list
	sections := make([]preview.Section, len(info.Sections))
	copy(sections, info.Sections)

	// Add a "Packages" section showing the actual package names
	pkgList := strings.Join(info.Packages, ", ")
	sections = append(sections, preview.Section{
		Label: "Includes",
		Value: pkgList,
	})

	p.Preview.SetSections(sections...)
}

func (p *Packages) Update(msg tea.KeyMsg) screens.Screen {
	p.Search.Update(msg)
	p.Menu.Update(msg)

	// Sync preview on every key (in case selection changed).
	p.syncPreview()

	switch msg.String() {

	case "esc":
		return screens.Home

	case "enter":
		// We'll hook package pages here later.
	}

	return screens.Packages
}

func (p *Packages) View(t theme.Theme) string {
	p.Header.SetRight(p.Search.View(t))

	// ── Panel dimensions ─────────────────────────────
	// Layout MaxWidth=94, padding=2 each side → inner=90
	// Left panel + 1 gap + right panel = 90
	leftW := 36
	rightW := 52
	panelH := 14

	leftBody := p.Menu.View(t)
	rightBody := p.Preview.View(t)

	left := card.New("Categories", leftBody, leftW).
		WithHeight(panelH).
		WithFocus(true).
		View(t)

	right := card.New("Preview", rightBody, rightW).
		WithHeight(panelH).
		View(t)

	// ── Status line between header and panels ────────
	sel := p.Menu.SelectedItem()
	posIndicator := lipgloss.NewStyle().
		Foreground(t.Palette.Muted).
		Render(fmt.Sprintf("%d/%d", p.Menu.Selected+1, len(p.Menu.Items)))

	statusLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(leftW+4).Render(posIndicator), // +4 for border
		lipgloss.NewStyle().
			Foreground(t.Palette.Overlay).
			Render(sel.Title),
	)

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		statusLine,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			left,
			" ",
			right,
		),
	)

	return layout.New(
		p.Header.View(t),
		body,
		p.Footer.View(t),
	).View()
}
