// Package preview provides a reusable preview panel.
package preview

import (
	"strings"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Preview struct {
	Title  string
	Body   string
	Meta   string
	Footer string
}

func New(title, body, meta, footer string) *Preview {
	return &Preview{
		Title:  title,
		Body:   body,
		Meta:   meta,
		Footer: footer,
	}
}

func (p *Preview) View(t theme.Theme) string {
	var lines []string

	if p.Title != "" {
		lines = append(lines,
			t.Styles.Title.Render(p.Title),
		)
	}

	if p.Body != "" {
		lines = append(lines, "")
		lines = append(lines,
			t.Styles.Normal.Render(p.Body),
		)
	}

	if p.Meta != "" {
		lines = append(lines, "")
		lines = append(lines,
			t.Styles.Muted.Render(p.Meta),
		)
	}

	if p.Footer != "" {
		lines = append(lines, "")
		lines = append(lines,
			t.Styles.Muted.Render(p.Footer),
		)
	}

	return strings.Join(lines, "\n")
}
