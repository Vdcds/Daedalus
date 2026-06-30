package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/pages"
)

func (a *App) View() string {
	var content string

	switch a.currentPage {

	case pages.Home:
		content = a.home.View()

	default:
		content = "Unknown Page"

	}

	return lipgloss.Place(
		a.windowWidth,
		a.windowHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
