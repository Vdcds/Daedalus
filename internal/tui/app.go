// Package tui gives a tui for the thing
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/pages/home"
	"github.com/vdcds/Daedalus/internal/tui/pages/packages"
)

type App struct {
	windowWidth  int
	windowHeight int

	currentPage pages.Page

	home     *home.Home
	packages *packages.Packages
}

func NewApp() *App {
	return &App{
		currentPage: pages.Home,
		home:        home.New(),
		packages:    packages.New(),
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}
