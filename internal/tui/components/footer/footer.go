// Package footer provides a reusable command bar.
package footer

import (
	"strings"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Action struct {
	Key         string
	Description string
}

type Footer struct {
	Actions []Action
}

func New(actions ...Action) *Footer {
	return &Footer{
		Actions: actions,
	}
}

func (f *Footer) View(t theme.Theme) string {
	var actions []string

	for _, action := range f.Actions {
		actions = append(
			actions,
			t.Styles.Accent.Render(action.Key)+
				" "+t.Styles.Help.Render(action.Description),
		)
	}

	return strings.Join(actions, "    ")
}
