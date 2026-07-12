package install

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case readyMsg:
		return m.waitForOutput()

	case lineMsg:
		if msg.Line != "" {
			m.Lines = append(m.Lines, msg.Line)

			const maxLines = 12

			if len(m.Lines) > maxLines {
				m.Lines = m.Lines[len(m.Lines)-maxLines:]
			}
		}

		return m.waitForOutput()

	case packageFinishedMsg:
		m.Completed = append(
			m.Completed,
			PackageResult{
				Name: msg.Package,
				Err:  msg.Err,
			},
		)

		m.Current++
		m.Lines = nil

		return m.startCurrent()

	default:
		var cmd tea.Cmd

		m.Spinner, cmd = m.Spinner.Update(msg)

		return cmd
	}
}
