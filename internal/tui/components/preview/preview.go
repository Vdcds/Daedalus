// Package preview provides an information pane with description
// text and aligned key-value metadata sections.
package preview

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Section is a key-value metadata row.
type Section struct {
	Label string
	Value string
}

// Preview renders descriptive text followed by aligned metadata.
type Preview struct {
	Title    string
	Content  []string
	Sections []Section
}

func New(title string, content ...string) *Preview {
	return &Preview{Title: title, Content: content}
}

func (p *Preview) SetTitle(title string)           { p.Title = title }
func (p *Preview) SetContent(content ...string)    { p.Content = content }
func (p *Preview) SetSections(sections ...Section) { p.Sections = sections }

func (p *Preview) View(t theme.Theme) string {
	var lines []string

	// ── Optional title (for standalone use outside a card) ──
	if p.Title != "" {
		lines = append(lines,
			lipgloss.NewStyle().Foreground(t.Palette.Text).Bold(true).Render(p.Title),
		)
		lines = append(lines, "")
	}

	// ── Description content ──
	for _, line := range p.Content {
		if line == "" {
			lines = append(lines, "")
		} else {
			lines = append(lines,
				lipgloss.NewStyle().Foreground(t.Palette.Muted).Render(line),
			)
		}
	}

	// ── Key-value metadata with column-aligned labels ──
	if len(p.Sections) > 0 {
		if len(lines) > 0 {
			lines = append(lines, "")
		}

		// Compute maximum label width for alignment.
		maxLabel := 0
		for _, s := range p.Sections {
			if len(s.Label) > maxLabel {
				maxLabel = len(s.Label)
			}
		}

		for _, s := range p.Sections {
			label := lipgloss.NewStyle().
				Foreground(t.Palette.Muted).
				Width(maxLabel).
				Render(s.Label)

			value := lipgloss.NewStyle().
				Foreground(t.Palette.Text).
				Render(s.Value)

			lines = append(lines, " "+label+"  "+value)
		}
	}

	return strings.Join(lines, "\n")
}
