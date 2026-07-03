// Package search provides a minimal inline search input.
package search

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Search is a simple text input with a "/" prefix.
type Search struct {
	Title       string
	Placeholder string
	Value       string
	Focused     bool
}

func New(placeholder string) *Search {
	return &Search{Placeholder: placeholder}
}

func (s *Search) Focus()       { s.Focused = true }
func (s *Search) Blur()        { s.Focused = false }
func (s *Search) Text() string { return s.Value }

func (s *Search) Update(msg tea.KeyMsg) {
	if !s.Focused {
		return
	}
	switch msg.Type {
	case tea.KeyBackspace:
		if len(s.Value) > 0 {
			s.Value = s.Value[:len(s.Value)-1]
		}
	case tea.KeyRunes:
		s.Value += string(msg.Runes)
	}
}

func (s *Search) View(t theme.Theme) string {
	prefix := lipgloss.NewStyle().
		Foreground(t.Palette.Pine).
		Bold(true).
		Render("/")

	value := s.Value
	if value == "" {
		value = lipgloss.NewStyle().
			Foreground(t.Palette.Overlay).
			Render(s.Placeholder)
	} else {
		value = lipgloss.NewStyle().
			Foreground(t.Palette.Text).
			Render(value)
	}

	return prefix + " " + value
}
