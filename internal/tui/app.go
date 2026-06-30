package tui

import tea "github.com/charmbracelet/bubbletea"

type App struct {
	width  int
	height int
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() tea.Cmd {
	return nil
}
