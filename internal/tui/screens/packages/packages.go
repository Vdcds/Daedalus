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

type Pane int

const (
	CategoryPane Pane = iota
	PackagePane
)

type Packages struct {
	Header *header.Header
	Search *search.Search

	CategoryList list.List
	PackageList  list.List
	FocusedPane  Pane

	Preview *preview.Preview
	Footer  *footer.Footer
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

func packageItems(category catalog.Category) []list.Item {
	items := make([]list.Item, 0, len(category.Packages))

	for _, pkg := range category.Packages {
		items = append(items, list.Item{
			ID:          pkg.BrewName,
			Title:       pkg.Name,
			Description: pkg.Description,
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
			Items:   menuItems(),
			Focused: true,
		},

		PackageList: list.List{
			Items: packageItems(catalog.Categories[0]),
		},

		FocusedPane: CategoryPane,

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
				Key:         "←→",
				Description: "Switch pane",
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

	switch msg.String() {
	case "esc":
		return screens.Home

	case "left":
		p.FocusedPane = CategoryPane

	case "right":
		p.FocusedPane = PackagePane

	case "up", "down":
		switch p.FocusedPane {
		case CategoryPane:
			before := p.CategoryList.Selected

			p.CategoryList.Update(msg)

			if before != p.CategoryList.Selected {
				selectedCategory := catalog.Categories[p.CategoryList.Selected]

				p.PackageList.Items = packageItems(selectedCategory)
				p.PackageList.Selected = 0
			}

		case PackagePane:
			p.PackageList.Update(msg)
		}

	case "enter":
		// Package actions come next.
	}

	return screens.Packages
}

func (p *Packages) View(t theme.Theme) string {
	p.Header.SetRight(p.Search.View(t))

	p.CategoryList.Focused = p.FocusedPane == CategoryPane
	p.PackageList.Focused = p.FocusedPane == PackagePane

	selectedCategory := catalog.Categories[p.CategoryList.Selected]
	selectedPackage := selectedCategory.Packages[p.PackageList.Selected]

	p.Preview.Title = selectedPackage.Name
	p.Preview.Body = selectedPackage.Description
	p.Preview.Meta = fmt.Sprintf(
		"brew install %s",
		selectedPackage.BrewName,
	)
	p.Preview.Footer = "Press Enter to install."

	left := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render("Categories"),
		t.Styles.Muted.Render(strings.Repeat("─", 18)),
		"",
		p.CategoryList.View(t),
	)

	middle := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render("Packages"),
		t.Styles.Muted.Render(strings.Repeat("─", 18)),
		"",
		p.PackageList.View(t),
	)

	right := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render("Preview"),
		t.Styles.Muted.Render(strings.Repeat("─", 18)),
		"",
		p.Preview.View(t),
	)

	divider := t.Styles.Muted.Render("│")

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().
			Width(28).
			Render(left),
		" "+divider+" ",
		lipgloss.NewStyle().
			Width(28).
			Render(middle),
		" "+divider+" ",
		lipgloss.NewStyle().
			Width(30).
			Render(right),
	)

	return layout.New(
		p.Header.View(t),
		body,
		p.Footer.View(t),
	).View()
}
