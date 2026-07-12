package install

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/vdcds/Daedalus/internal/installer"
)

type Job struct {
	Name     string
	BrewName string
}

type lineMsg struct {
	Line string
}

type packageFinishedMsg struct {
	Job Job
	Err error
}

type readyMsg struct{}

type FinishedMsg struct {
	Err error
}

type PackageResult struct {
	Job Job
	Err error
}

type Model struct {
	Spinner spinner.Model
	Lines   []string

	Jobs      []Job
	Current   int
	Completed []PackageResult

	runner installer.Runner

	lines chan string
	done  chan error
}

func New() *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &Model{
		Spinner:   s,
		Lines:     make([]string, 0, 32),
		Completed: make([]PackageResult, 0),
	}
}

func (m *Model) Start(
	runner installer.Runner,
	jobs []Job,
) tea.Cmd {
	m.runner = runner
	m.Jobs = append([]Job(nil), jobs...)
	m.Current = 0
	m.Completed = nil
	m.Lines = nil

	return m.startCurrent()
}

func (m *Model) startCurrent() tea.Cmd {
	if m.Current >= len(m.Jobs) {
		return m.finish()
	}

	job := m.Jobs[m.Current]

	return func() tea.Msg {
		process, err := m.runner.Start(
			[]string{job.BrewName},
		)

		if err != nil {
			return packageFinishedMsg{
				Job: job,
				Err: err,
			}
		}

		m.lines = make(chan string, 64)
		m.done = make(chan error, 1)

		go m.readProcess(process)

		return readyMsg{}
	}
}

func (m *Model) readProcess(process installer.Process) {
	var wg sync.WaitGroup

	wg.Add(2)

	go scan(
		&wg,
		process.Stdout(),
		m.lines,
	)

	go scan(
		&wg,
		process.Stderr(),
		m.lines,
	)

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
	job := m.Jobs[m.Current]

	return func() tea.Msg {
		line, ok := <-m.lines

		if ok {
			return lineMsg{
				Line: line,
			}
		}

		return packageFinishedMsg{
			Job: job,
			Err: <-m.done,
		}
	}
}

func (m *Model) finish() tea.Cmd {
	return func() tea.Msg {
		failed := 0

		for _, result := range m.Completed {
			if result.Err != nil {
				failed++
			}
		}

		if failed > 0 {
			return FinishedMsg{
				Err: fmt.Errorf(
					"%d package(s) failed to install",
					failed,
				),
			}
		}

		return FinishedMsg{}
	}
}
