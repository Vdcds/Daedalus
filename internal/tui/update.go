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

		case "ctrl+c", "q":
			return a, tea.Quit
		}

		switch a.currentScreen {

		case screens.Home:
			a.currentScreen = a.home.Update(msg)

		case screens.Packages:
			// Packages screen update will come later.
		}
	}

	return a, nil
}
