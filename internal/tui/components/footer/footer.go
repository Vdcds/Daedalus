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
	divider := t.Styles.Muted.Render(
		strings.Repeat("─", 92),
	)

	var actions []string

	for _, action := range f.Actions {
		actions = append(
			actions,
			t.Styles.Highlight.Render(action.Key)+" "+
				t.Styles.Muted.Render(action.Description),
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,

		divider,

		"",

		strings.Join(actions, "   ·   "),
	)
}
