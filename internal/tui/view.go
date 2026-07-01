package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/screens"
)

func (a *App) View() string {
	var content string

	switch a.currentScreen {

	case screens.Home:
		content = a.home.View(a.theme)

	case screens.Packages:
		content = a.packages.View()

	default:
		content = "Unknown Screen"
	}

	return lipgloss.Place(
		a.windowWidth,
		a.windowHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
