package install

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case lineMsg:
		if msg.Line != "" {
			m.Lines = append(m.Lines, msg.Line)

			const maxLines = 14

			if len(m.Lines) > maxLines {
				m.Lines = m.Lines[len(m.Lines)-maxLines:]
			}
		}

		return m.waitForOutput()

	case finishedMsg:
		return func() tea.Msg {
			return FinishedMsg{
				Err: msg.Err,
			}
		}

	default:
		var cmd tea.Cmd

		m.Spinner, cmd = m.Spinner.Update(msg)

		return cmd
	}
}

type FinishedMsg struct {
	Err error
}
