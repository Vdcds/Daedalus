package installer

import (
	"fmt"
	"os/exec"
	"strings"
)

type Homebrew struct{}

func NewHomebrew() *Homebrew {
	return &Homebrew{}
}

func (h *Homebrew) Install(packages []string) Result {
	if len(packages) == 0 {
		return Result{
			Err: fmt.Errorf("no packages selected"),
		}
	}

	args := append([]string{"install"}, packages...)

	cmd := exec.Command("brew", args...)

	output, err := cmd.CombinedOutput()

	return Result{
		Output: strings.TrimSpace(string(output)),
		Err:    err,
	}
}
