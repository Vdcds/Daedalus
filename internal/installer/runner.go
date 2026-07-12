// Package installer provides software installation backends.
package installer

import "io"

type Process interface {
	Stdout() io.ReadCloser
	Stderr() io.ReadCloser
	Wait() error
}

type Runner interface {
	Start(packages []string) (Process, error)
}
