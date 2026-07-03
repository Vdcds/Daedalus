// Package menu provides a reusable selectable menu component.
package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

// Item is a single menu entry.
type Item struct {
	ID          string
	Title       string
	Description string
	Icon        string
	Badge       string // Right-aligned label (e.g. count).
}

// Menu is a navigable list of items.
type Menu struct {
	Items    []Item
	Selected int
	Width    int // If > 0, compact single-line mode with right-aligned badges.
}

func (m *Menu) Update(msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		if m.Selected > 0 {
			m.Selected--
		}
	case "down":
		if m.Selected < len(m.Items)-1 {
			m.Selected++
		}
	}
}

// SelectedItem returns the currently highlighted item.
func (m Menu) SelectedItem() Item {
	return m.Items[m.Selected]
}

func (m Menu) View(t theme.Theme) string {
	var lines []string
	for i, item := range m.Items {
		lines = append(lines, m.renderItem(item, i == m.Selected, t))
	}
	if m.Width > 0 {
		return strings.Join(lines, "\n")
	}
	// Extra spacing between items in descriptive mode.
	return strings.Join(lines, "\n\n")
}

func (m Menu) renderItem(item Item, selected bool, t theme.Theme) string {
	// Cursor — a solid vertical bar for the active item.
	cursor := "  "
	if selected {
		cursor = lipgloss.NewStyle().
			Foreground(t.Palette.Pine).
			Bold(true).
			Render("▎ ")
	}

	// Title styling.
	var title string
	if selected {
		title = lipgloss.NewStyle().
			Foreground(t.Palette.Text).
			Bold(true).
			Render(item.Title)
	} else {
		title = lipgloss.NewStyle().
			Foreground(t.Palette.Muted).
			Render(item.Title)
	}

	// ── Compact mode (Width > 0): single line + right-aligned badge ──
	if m.Width > 0 {
		line := cursor + title

		if item.Badge != "" {
			var badge string
			if selected {
				badge = lipgloss.NewStyle().
					Foreground(t.Palette.Pine).
					Render(item.Badge)
			} else {
				badge = lipgloss.NewStyle().
					Foreground(t.Palette.Overlay).
					Render(item.Badge)
			}

			usedW := lipgloss.Width(cursor) + lipgloss.Width(title)
			badgeW := lipgloss.Width(badge)
			pad := m.Width - usedW - badgeW
			if pad < 1 {
				pad = 1
			}
			line = cursor + title + strings.Repeat(" ", pad) + badge
		}

		return line
	}

	// ── Descriptive mode (Width = 0): title + description below ──
	if item.Description != "" {
		desc := lipgloss.NewStyle().
			Foreground(t.Palette.Muted).
			Render("  " + item.Description)
		return cursor + title + "\n" + desc
	}

	return cursor + title
}
