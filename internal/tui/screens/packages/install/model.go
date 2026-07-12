package install

import (
	"bufio"
	"io"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/installer"
)

type lineMsg struct {
	Line string
}

type finishedMsg struct {
	Err error
}

type Model struct {
	Spinner spinner.Model

	Lines []string

	process installer.Process

	lines chan string
	done  chan error

	started bool
}

func New() *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &Model{
		Spinner: s,
		Lines:   make([]string, 0, 32),
	}
}

func (m *Model) Start(
	runner installer.Runner,
	packages []string,
) tea.Cmd {
	return func() tea.Msg {
		process, err := runner.Start(packages)
		if err != nil {
			return finishedMsg{
				Err: err,
			}
		}

		m.process = process
		m.lines = make(chan string, 64)
		m.done = make(chan error, 1)
		m.started = true

		go m.readProcess(process)

		return lineMsg{}
	}
}

func (m *Model) readProcess(process installer.Process) {
	var wg sync.WaitGroup

	wg.Add(2)

	go scan(&wg, process.Stdout(), m.lines)
	go scan(&wg, process.Stderr(), m.lines)

	go func() {
		wg.Wait()

		err := process.Wait()

		close(m.lines)

		m.done <- err
		close(m.done)
	}()
}

func scan(
	wg *sync.WaitGroup,
	reader io.Reader,
	lines chan<- string,
) {
	defer wg.Done()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines <- scanner.Text()
	}
}

func (m *Model) waitForOutput() tea.Cmd {
	return func() tea.Msg {
		line, ok := <-m.lines

		if ok {
			return lineMsg{
				Line: line,
			}
		}

		return finishedMsg{
			Err: <-m.done,
		}
	}
}
