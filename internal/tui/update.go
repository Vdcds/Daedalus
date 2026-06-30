package tui

import (
	"github.com/vdcds/Daedalus/internal/tui/pages"

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
				a.home.Menu.MoveUp()
			}

		case "down":
			switch a.currentPage {
			case pages.Home:
				a.home.Menu.MoveDown()
			}
		}
	}

	return a, nil
}
