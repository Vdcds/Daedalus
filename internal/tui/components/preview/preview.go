// Package preview provides a reusable information preview component.
package preview

import (
	"strings"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Preview struct {
	Title   string
	Content []string
}

func New(title string, content ...string) *Preview {
	return &Preview{
		Title:   title,
		Content: content,
	}
}

func (p *Preview) SetTitle(title string) {
	p.Title = title
}

func (p *Preview) SetContent(content ...string) {
	p.Content = content
}

func (p *Preview) View(t theme.Theme) string {
	var lines []string

	if p.Title != "" {
		lines = append(lines, t.Styles.Selected.Render(p.Title))
		lines = append(lines, "")
	}

	for _, line := range p.Content {
		lines = append(lines, t.Styles.Normal.Render(line))
	}

	return strings.Join(lines, "\n")
}
