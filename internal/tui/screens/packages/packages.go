// Package packages renders the packages screen.
package packages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/catalog"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	list "github.com/vdcds/Daedalus/internal/tui/components/list"
	"github.com/vdcds/Daedalus/internal/tui/components/preview"
	"github.com/vdcds/Daedalus/internal/tui/components/search"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Packages struct {
	Header       *header.Header
	Search       *search.Search
	CategoryList list.List
	Preview      *preview.Preview
	Footer       *footer.Footer
}

func menuItems() []list.Item {
	items := make([]list.Item, 0, len(catalog.Categories))

	for _, category := range catalog.Categories {
		items = append(items, list.Item{
			ID:          category.ID,
			Title:       category.Name,
			Description: category.Description,
		})
	}

	return items
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

		CategoryList: list.List{
			Items: menuItems(),
		},

		Preview: preview.New(
			"",
			"",
			"",
			"",
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
	p.CategoryList.Update(msg)

	switch msg.String() {

	case "esc":
		return screens.Home

	case "enter":
		// Category navigation comes later.
	}

	return screens.Packages
}

func (p *Packages) View(t theme.Theme) string {
	p.Header.SetRight(p.Search.View(t))

	selected := catalog.Categories[p.CategoryList.Selected]

	p.Preview.Title = selected.Name
	p.Preview.Body = selected.Description
	p.Preview.Meta = fmt.Sprintf(
		"%d packages",
		len(selected.Packages),
	)
	p.Preview.Footer = "Press Enter to browse packages."

	left := lipgloss.JoinVertical(
		lipgloss.Left,

		t.Styles.Highlight.Render("Categories"),
		t.Styles.Muted.Render(strings.Repeat("─", 18)),
		"",
		p.CategoryList.View(t),
	)

	right := lipgloss.JoinVertical(
		lipgloss.Left,

		t.Styles.Highlight.Render("Preview"),
		t.Styles.Muted.Render(strings.Repeat("─", 18)),
		"",
		p.Preview.View(t),
	)

	body := layout.NewSplit(
		left,
		right,
	).View(t)

	return layout.New(
		p.Header.View(t),
		body,
		p.Footer.View(t),
	).View()
}
