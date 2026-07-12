// Package packages renders the packages screen.
package packages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/catalog"
	"github.com/vdcds/Daedalus/internal/installer"

	"github.com/vdcds/Daedalus/internal/tui/components/footer"
	"github.com/vdcds/Daedalus/internal/tui/components/header"
	layout "github.com/vdcds/Daedalus/internal/tui/components/layout"
	list "github.com/vdcds/Daedalus/internal/tui/components/list"
	"github.com/vdcds/Daedalus/internal/tui/components/preview"
	"github.com/vdcds/Daedalus/internal/tui/components/search"
	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/screens/packages/install"
	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Pane int

const (
	CategoryPane Pane = iota
	PackagePane
)

type Mode int

const (
	BrowseMode Mode = iota
	ConfirmMode
	InstallingMode
	ResultMode
)

type Packages struct {
	Header *header.Header
	Search *search.Search

	CategoryList list.List
	PackageList  list.List
	FocusedPane  Pane

	Selected map[string]bool
	Mode     Mode

	Install    *install.Model
	InstallErr error
	Installer  installer.Runner

	Preview *preview.Preview
	Footer  *footer.Footer
}

func menuItems() []list.Item {
	items := make(
		[]list.Item,
		0,
		len(catalog.Categories),
	)

	for _, category := range catalog.Categories {
		items = append(
			items,
			list.Item{
				ID:          category.ID,
				Title:       category.Name,
				Description: category.Description,
			},
		)
	}

	return items
}

func packageItems(
	category catalog.Category,
	selected map[string]bool,
) []list.Item {
	items := make(
		[]list.Item,
		0,
		len(category.Packages),
	)

	for _, pkg := range category.Packages {
		items = append(
			items,
			list.Item{
				ID:          pkg.BrewName,
				Title:       pkg.Name,
				Description: pkg.Description,
				Marked:      selected[pkg.BrewName],
			},
		)
	}

	return items
}

func paneHeading(
	title string,
	focused bool,
	t theme.Theme,
) string {
	titleStyle := t.Styles.Muted
	dividerStyle := t.Styles.Muted

	if focused {
		titleStyle = t.Styles.Highlight
		dividerStyle = t.Styles.Highlight
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(title),
		dividerStyle.Render(
			strings.Repeat("─", 18),
		),
	)
}

func New() *Packages {
	s := search.New("Search packages...")
	s.Focus()

	selected := make(map[string]bool)

	return &Packages{
		Header: header.New(
			"📦 Packages",
			"Browse package categories",
		),

		Search: s,

		CategoryList: list.List{
			Items: menuItems(),
		},

		PackageList: list.List{
			Items: packageItems(
				catalog.Categories[0],
				selected,
			),
		},

		FocusedPane: CategoryPane,

		Selected: selected,
		Mode:     BrowseMode,

		Install:   install.New(),
		Installer: installer.NewHomebrew(),

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
				Key:         "Space",
				Description: "Select",
			},
			footer.Action{
				Key:         "Enter",
				Description: "Install",
			},
			footer.Action{
				Key:         "Esc",
				Description: "Back",
			},
		),
	}
}

func (p *Packages) Update(
	msg tea.Msg,
) (screens.Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case install.FinishedMsg:
		p.InstallErr = msg.Err
		p.Mode = ResultMode

		return screens.Packages, nil
	}

	if p.Mode == InstallingMode {
		cmd := p.Install.Update(msg)

		return screens.Packages, cmd
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		return p.updateKey(keyMsg)
	}

	return screens.Packages, nil
}

func (p *Packages) updateKey(
	msg tea.KeyMsg,
) (screens.Screen, tea.Cmd) {
	switch p.Mode {
	case ConfirmMode:
		return p.updateConfirm(msg)

	case ResultMode:
		return p.updateResult(msg)

	default:
		return p.updateBrowser(msg)
	}
}

func (p *Packages) updateBrowser(
	msg tea.KeyMsg,
) (screens.Screen, tea.Cmd) {
	p.Search.Update(msg)

	switch msg.String() {
	case "esc":
		return screens.Home, nil

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
				selectedCategory :=
					catalog.Categories[p.CategoryList.Selected]

				p.PackageList.Items = packageItems(
					selectedCategory,
					p.Selected,
				)

				p.PackageList.Selected = 0
			}

		case PackagePane:
			p.PackageList.Update(msg)
		}

	case " ":
		if p.FocusedPane == PackagePane {
			p.toggleSelectedPackage()
		}

	case "enter":
		if selectedCount(p.Selected) > 0 {
			p.Mode = ConfirmMode
		}
	}

	return screens.Packages, nil
}

func (p *Packages) updateConfirm(
	msg tea.KeyMsg,
) (screens.Screen, tea.Cmd) {
	switch msg.String() {
	case "esc":
		p.Mode = BrowseMode

	case "enter":
		p.Mode = InstallingMode
		p.Install = install.New()
		p.InstallErr = nil

		jobs := p.selectedJobs()

		return screens.Packages, tea.Batch(
			p.Install.Spinner.Tick,
			p.Install.Start(
				p.Installer,
				jobs,
			),
		)
	}

	return screens.Packages, nil
}

func (p *Packages) updateResult(
	msg tea.KeyMsg,
) (screens.Screen, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc":
		if p.InstallErr == nil {
			p.clearSelection()
		}

		p.Mode = BrowseMode
	}

	return screens.Packages, nil
}

func (p *Packages) toggleSelectedPackage() {
	item := p.PackageList.SelectedItem()

	if item.ID == "" {
		return
	}

	p.Selected[item.ID] = !p.Selected[item.ID]

	selectedCategory :=
		catalog.Categories[p.CategoryList.Selected]

	currentSelection := p.PackageList.Selected

	p.PackageList.Items = packageItems(
		selectedCategory,
		p.Selected,
	)

	p.PackageList.Selected = currentSelection
}

func (p *Packages) clearSelection() {
	p.Selected = make(map[string]bool)

	selectedCategory :=
		catalog.Categories[p.CategoryList.Selected]

	p.PackageList.Items = packageItems(
		selectedCategory,
		p.Selected,
	)
}

func (p *Packages) View(
	t theme.Theme,
) string {
	switch p.Mode {
	case ConfirmMode:
		return p.confirmView(t)

	case InstallingMode:
		return p.installingView(t)

	case ResultMode:
		return p.resultView(t)

	default:
		return p.browserView(t)
	}
}

func (p *Packages) browserView(
	t theme.Theme,
) string {
	p.Header.SetRight(
		p.Search.View(t),
	)

	selectedCategory :=
		catalog.Categories[p.CategoryList.Selected]

	selectedPackage :=
		selectedCategory.Packages[p.PackageList.Selected]

	p.Preview.Title = selectedPackage.Name
	p.Preview.Body = selectedPackage.Description

	if p.Selected[selectedPackage.BrewName] {
		p.Preview.Meta =
			"✓ Selected for installation"
	} else {
		p.Preview.Meta = fmt.Sprintf(
			"brew install %s",
			selectedPackage.BrewName,
		)
	}

	p.Preview.Footer = fmt.Sprintf(
		"%d selected · Space to toggle",
		selectedCount(p.Selected),
	)

	left := lipgloss.JoinVertical(
		lipgloss.Left,
		paneHeading(
			"Categories",
			p.FocusedPane == CategoryPane,
			t,
		),
		"",
		p.CategoryList.View(t),
	)

	middle := lipgloss.JoinVertical(
		lipgloss.Left,
		paneHeading(
			"Packages",
			p.FocusedPane == PackagePane,
			t,
		),
		"",
		p.PackageList.View(t),
	)

	right := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Muted.Render("Preview"),
		t.Styles.Muted.Render(
			strings.Repeat("─", 18),
		),
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

func (p *Packages) confirmView(
	t theme.Theme,
) string {
	count := selectedCount(p.Selected)

	pageHeader := header.New(
		"📦 Confirm Installation",
		fmt.Sprintf(
			"%d package%s selected",
			count,
			plural(count),
		),
	)

	var packages []string

	for _, category := range catalog.Categories {
		for _, pkg := range category.Packages {
			if !p.Selected[pkg.BrewName] {
				continue
			}

			packages = append(
				packages,
				t.Styles.Highlight.Render("✓")+
					" "+
					t.Styles.Title.Render(pkg.Name)+
					"\n"+
					"  "+
					t.Styles.Muted.Render(
						pkg.Description,
					),
			)
		}
	}

	queue := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render(
			"Install Queue",
		),
		t.Styles.Muted.Render(
			strings.Repeat("─", 42),
		),
		"",
		strings.Join(packages, "\n\n"),
	)

	command := fmt.Sprintf(
		"brew install %s",
		strings.Join(
			p.selectedBrewNames(),
			" ",
		),
	)

	summary := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Muted.Render("Command"),
		"",
		t.Styles.Normal.Render(command),
		"",
		t.Styles.Muted.Render(
			"Review the selected packages before installation.",
		),
	)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().
			Width(48).
			Render(queue),
		" "+
			t.Styles.Muted.Render("│")+
			" ",
		lipgloss.NewStyle().
			Width(40).
			Render(summary),
	)

	confirmFooter := footer.New(
		footer.Action{
			Key:         "Enter",
			Description: "Confirm",
		},
		footer.Action{
			Key:         "Esc",
			Description: "Cancel",
		},
	)

	return layout.New(
		pageHeader.View(t),
		body,
		confirmFooter.View(t),
	).View()
}

func (p *Packages) installingView(
	t theme.Theme,
) string {
	count := selectedCount(p.Selected)

	pageHeader := header.New(
		"📦 Installing",
		fmt.Sprintf(
			"Installing %d package%s",
			count,
			plural(count),
		),
	)

	installFooter := footer.New(
		footer.Action{
			Key:         "Ctrl+C",
			Description: "Quit",
		},
	)

	return layout.New(
		pageHeader.View(t),
		p.Install.View(t),
		installFooter.View(t),
	).View()
}

func (p *Packages) resultView(
	t theme.Theme,
) string {
	success := p.InstallErr == nil

	title := "✓ Installation Complete"
	subtitle :=
		"Selected packages installed successfully"

	if !success {
		title = "Installation Failed"
		subtitle = p.InstallErr.Error()
	}

	pageHeader := header.New(
		title,
		subtitle,
	)

	output := "Installation finished."

	if len(p.Install.Lines) > 0 {
		output = strings.Join(
			p.Install.Lines,
			"\n",
		)
	}

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		t.Styles.Highlight.Render(
			"Homebrew Output",
		),
		t.Styles.Muted.Render(
			strings.Repeat("─", 72),
		),
		"",
		t.Styles.Normal.Render(output),
	)

	resultFooter := footer.New(
		footer.Action{
			Key:         "Enter",
			Description: "Continue",
		},
		footer.Action{
			Key:         "Esc",
			Description: "Back",
		},
	)

	return layout.New(
		pageHeader.View(t),
		body,
		resultFooter.View(t),
	).View()
}

func (p *Packages) selectedJobs() []install.Job {
	var jobs []install.Job

	for _, category := range catalog.Categories {
		for _, pkg := range category.Packages {
			if !p.Selected[pkg.BrewName] {
				continue
			}

			jobs = append(
				jobs,
				install.Job{
					Name:     pkg.Name,
					BrewName: pkg.BrewName,
				},
			)
		}
	}

	return jobs
}

func (p *Packages) selectedBrewNames() []string {
	var names []string

	for _, job := range p.selectedJobs() {
		names = append(
			names,
			job.BrewName,
		)
	}

	return names
}

func selectedCount(
	selected map[string]bool,
) int {
	count := 0

	for _, marked := range selected {
		if marked {
			count++
		}
	}

	return count
}

func plural(count int) string {
	if count == 1 {
		return ""
	}

	return "s"
}
