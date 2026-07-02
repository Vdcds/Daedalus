// Package search provides a reusable search input component.
package search

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Search struct {
	Title       string
	Placeholder string
	Value       string
	Focused     bool
}

func New(placeholder string) *Search {
	return &Search{
		Placeholder: placeholder,
	}
}

func (s *Search) Focus() {
	s.Focused = true
}

func (s *Search) Blur() {
	s.Focused = false
}

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

func (s *Search) Text() string {
	return s.Value
}

func (s *Search) View(t theme.Theme) string {
	value := s.Value

	if value == "" {
		value = t.Styles.Muted.Render(s.Placeholder)
	}

	prefix := "  "

	if s.Focused {
		prefix = "❯ "
	}

	return t.Styles.Normal.Render(
		fmt.Sprintf("%s%s", prefix, value),
	)
}
