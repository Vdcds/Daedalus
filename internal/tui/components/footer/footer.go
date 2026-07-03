// Package footer provides a compact keybinding hint bar.
package footer

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Action is a single key hint.
type Action struct {
	Key         string
	Description string
}

// Footer renders a horizontal list of key hints.
type Footer struct {
	Actions []Action
}

func New(actions ...Action) *Footer {
	return &Footer{Actions: actions}
}

func (f *Footer) View(t theme.Theme) string {
	var parts []string

	for _, a := range f.Actions {
		key := lipgloss.NewStyle().
			Foreground(t.Palette.Pine).
			Bold(true).
			Render(a.Key)
		desc := lipgloss.NewStyle().
			Foreground(t.Palette.Muted).
			Render(" " + a.Description)
		parts = append(parts, key+desc)
	}

	sep := lipgloss.NewStyle().
		Foreground(t.Palette.Overlay).
		Render("  ·  ")

	return strings.Join(parts, sep)
}
