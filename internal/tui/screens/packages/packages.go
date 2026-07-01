// Package packages renders the Packages page.
package packages

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle    = lipgloss.NewStyle().Bold(true)
	subtitleStyle = lipgloss.NewStyle().Faint(true)
)

type Packages struct{}

func New() *Packages {
	return &Packages{}
}

func (p *Packages) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,

		titleStyle.Render("📦 Packages"),

		"",

		subtitleStyle.Render("Package registry coming soon."),

		"",

		subtitleStyle.Render("(press esc to go back)"),
	)
}
