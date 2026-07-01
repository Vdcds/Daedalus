// Package header provides a reusable page header.
package header

import (
	"strings"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

type Header struct {
	Title    string
	Subtitle string
}

func New(title, subtitle string) *Header {
	return &Header{
		Title:    title,
		Subtitle: subtitle,
	}
}

func (h *Header) View(t theme.Theme) string {
	var lines []string

	if h.Title != "" {
		lines = append(lines,
			t.Styles.Title.Render(h.Title),
		)
	}

	if h.Subtitle != "" {
		lines = append(lines,
			t.Styles.Subtitle.Render(h.Subtitle),
		)
	}

	return strings.Join(lines, "\n")
}
