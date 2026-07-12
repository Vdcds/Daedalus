package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/screens"
)

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.windowWidth = msg.Width
		a.windowHeight = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit

		case "q":
			if a.currentScreen == screens.Home {
				return a, tea.Quit
			}
		}
	}

	switch a.currentScreen {
	case screens.Home:
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			a.currentScreen = a.home.Update(keyMsg)
		}

	case screens.Packages:
		nextScreen, cmd := a.packages.Update(msg)
		a.currentScreen = nextScreen

		return a, cmd
	}

	return a, nil
}
