package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/screens"
	"github.com/vdcds/Daedalus/internal/tui/screens/home"
	"github.com/vdcds/Daedalus/internal/tui/screens/packages"
	"github.com/vdcds/Daedalus/internal/tui/theme"
	"github.com/vdcds/Daedalus/internal/tui/themes/rosepine"
)

type App struct {
	windowWidth  int
	windowHeight int

	currentScreen screens.Screen

	theme theme.Theme

	home     *home.Home
	packages *packages.Packages
}

func NewApp() *App {
	return &App{
		currentScreen: screens.Home,

		theme: rosepine.New(),

		home:     home.New(),
		packages: packages.New(),
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}
