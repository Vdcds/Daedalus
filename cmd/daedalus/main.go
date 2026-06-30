package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui"
)

func main() {
	p := tea.NewProgram(
		tui.NewApp(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
