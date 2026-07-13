package installer

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Homebrew struct{}

type homebrewProcess struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stderr io.ReadCloser
}

func NewHomebrew() *Homebrew {
	return &Homebrew{}
}

func (h *Homebrew) Start(packages []string) (Process, error) {
	if len(packages) == 0 {
		return nil, fmt.Errorf("no packages selected")
	}

	args := append([]string{"install"}, packages...)

	cmd := exec.Command("brew", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start homebrew: %w", err)
	}

	return &homebrewProcess{
		cmd:    cmd,
		stdout: stdout,
		stderr: stderr,
	}, nil
}

func (h *Homebrew) Installed() (map[string]bool, error) {
	installed := make(map[string]bool)

	if err := collectInstalled(
		installed,
		"list",
		"--formula",
		"-1",
	); err != nil {
		return nil, fmt.Errorf(
			"list installed formulae: %w",
			err,
		)
	}

	if err := collectInstalled(
		installed,
		"list",
		"--cask",
		"-1",
	); err != nil {
		return nil, fmt.Errorf(
			"list installed casks: %w",
			err,
		)
	}

	return installed, nil
}

func collectInstalled(
	installed map[string]bool,
	args ...string,
) error {
	cmd := exec.Command("brew", args...)

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	for _, name := range strings.Fields(
		string(output),
	) {
		installed[name] = true
	}

	return nil
}

func (p *homebrewProcess) Stdout() io.ReadCloser {
	return p.stdout
}

func (p *homebrewProcess) Stderr() io.ReadCloser {
	return p.stderr
}

func (p *homebrewProcess) Wait() error {
	return p.cmd.Wait()
}
