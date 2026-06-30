// Package tui is a terminal-ui-parts toolkit
package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/pages/home"
)

func (a *App) View() string {
	var content string

	switch a.currentPage {

	case pages.Home:
		content = home.View()

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
