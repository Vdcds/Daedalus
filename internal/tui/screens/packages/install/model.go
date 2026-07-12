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

type lineMsg struct {
	Line string
}

type packageFinishedMsg struct {
	Package string
	Err     error
}

type readyMsg struct{}

type FinishedMsg struct {
	Err error
}

type PackageResult struct {
	Name string
	Err  error
}

type Model struct {
	Spinner spinner.Model
	Lines   []string

	Packages  []string
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
	packages []string,
) tea.Cmd {
	m.runner = runner
	m.Packages = append([]string(nil), packages...)
	m.Current = 0
	m.Completed = nil
	m.Lines = nil

	return m.startCurrent()
}

func (m *Model) startCurrent() tea.Cmd {
	if m.Current >= len(m.Packages) {
		return m.finish()
	}

	pkg := m.Packages[m.Current]

	return func() tea.Msg {
		process, err := m.runner.Start([]string{pkg})
		if err != nil {
			return packageFinishedMsg{
				Package: pkg,
				Err:     err,
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
	pkg := m.Packages[m.Current]

	return func() tea.Msg {
		line, ok := <-m.lines

		if ok {
			return lineMsg{
				Line: line,
			}
		}

		return packageFinishedMsg{
			Package: pkg,
			Err:     <-m.done,
		}
	}
}

func (m *Model) finish() tea.Cmd {
	return func() tea.Msg {
		var failed []string

		for _, result := range m.Completed {
			if result.Err != nil {
				failed = append(failed, result.Name)
			}
		}

		if len(failed) > 0 {
			return FinishedMsg{
				Err: fmt.Errorf(
					"%d package(s) failed to install",
					len(failed),
				),
			}
		}

		return FinishedMsg{}
	}
}
