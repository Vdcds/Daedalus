package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/pages"
)

type App struct {
	windowWidth  int
	windowHeight int

	currentPage pages.Page
}

func NewApp() *App {
	return &App{
		currentPage: pages.Home,
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}
