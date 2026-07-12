package installer

import (
	"fmt"
	"io"
	"os/exec"
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

func (p *homebrewProcess) Stdout() io.ReadCloser {
	return p.stdout
}

func (p *homebrewProcess) Stderr() io.ReadCloser {
	return p.stderr
}

func (p *homebrewProcess) Wait() error {
	return p.cmd.Wait()
}
