// Package tui renders terminal UI
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/pages/home"
)

type App struct {
	windowWidth  int
	windowHeight int

	currentPage pages.Page

	home *home.Home
}

func NewApp() *App {
	return &App{
		currentPage: pages.Home,
		home:        home.New(),
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}
