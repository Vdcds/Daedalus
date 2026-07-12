package install

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/vdcds/Daedalus/internal/tui/theme"
)

func (m *Model) View(t theme.Theme) string {
	currentPackage := "Preparing installation"

	if m.Current < len(m.Jobs) {
		currentPackage = m.Jobs[m.Current].Name
	}

	status := lipgloss.JoinHorizontal(
		lipgloss.Left,
		t.Styles.Highlight.Render(
			m.Spinner.View(),
		),
		" ",
		t.Styles.Title.Render(
			fmt.Sprintf(
				"Installing %s",
				currentPackage,
			),
		),
	)

	progress := t.Styles.Muted.Render(
		fmt.Sprintf(
			"%d of %d packages completed",
			len(m.Completed),
			len(m.Jobs),
		),
	)

	results := m.resultsView(t)

	logs := "Waiting for Homebrew output..."

	if len(m.Lines) > 0 {
		logs = strings.Join(
			m.Lines,
			"\n",
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		status,
		progress,
		"",
		results,
		"",
		t.Styles.Muted.Render(
			strings.Repeat("─", 72),
		),
		"",
		t.Styles.Normal.Render(logs),
	)
}

func (m *Model) resultsView(
	t theme.Theme,
) string {
	if len(m.Completed) == 0 {
		return t.Styles.Muted.Render(
			"No packages completed yet.",
		)
	}

	var rows []string

	for _, result := range m.Completed {
		if result.Err != nil {
			rows = append(
				rows,
				t.Styles.Muted.Render("✗")+" "+
					t.Styles.Normal.Render(
						result.Job.Name,
					),
			)

			continue
		}

		rows = append(
			rows,
			t.Styles.Highlight.Render("✓")+" "+
				t.Styles.Normal.Render(
					result.Job.Name,
				),
		)
	}

	return strings.Join(rows, "\n")
}
