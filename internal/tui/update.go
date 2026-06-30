package tui

import (
	"github.com/vdcds/Daedalus/internal/tui/pages"
	"github.com/vdcds/Daedalus/internal/tui/pages/home"

	tea "github.com/charmbracelet/bubbletea"
)

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		a.windowWidth = msg.Width
		a.windowHeight = msg.Height

	case tea.KeyMsg:

		switch msg.String() {

		case "q", "ctrl+c":
			return a, tea.Quit

		case "up":

			switch a.currentPage {
			case pages.Home:
				home.MainMenu.MoveUp()
			}

		case "down":

			switch a.currentPage {
			case pages.Home:
				home.MainMenu.MoveDown()
			}
		}
	}

	return a, nil
}
