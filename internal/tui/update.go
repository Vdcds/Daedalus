package tui

import (
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

		default:
			a.home.Update(msg)

		}
	}

	return a, nil
}
