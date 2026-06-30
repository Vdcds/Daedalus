package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/pages/home"
	"github.com/vdcds/Daedalus/internal/tui/pages/packages"
	"github.com/vdcds/Daedalus/internal/tui/theme"
	"github.com/vdcds/Daedalus/internal/tui/themes/rosepine"
)

type App struct {
	windowWidth  int
	windowHeight int

	currentPage pages.Page

	theme theme.Theme

	home     *home.Home
	packages *packages.Packages
}

func NewApp() *App {
	return &App{
		currentPage: pages.Home,

		theme: rosepine.New(),

		home:     home.New(),
		packages: packages.New(),
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}
