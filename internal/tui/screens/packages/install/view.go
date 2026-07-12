package install

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

func (m *Model) View(t theme.Theme) string {
	status := lipgloss.JoinHorizontal(
		lipgloss.Left,
		t.Styles.Highlight.Render(m.Spinner.View()),
		" ",
		t.Styles.Title.Render("Homebrew is installing packages"),
	)

	logs := "Waiting for Homebrew output..."

	if len(m.Lines) > 0 {
		logs = strings.Join(m.Lines, "\n")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		status,
		"",
		t.Styles.Muted.Render(strings.Repeat("─", 72)),
		"",
		t.Styles.Normal.Render(logs),
	)
}
