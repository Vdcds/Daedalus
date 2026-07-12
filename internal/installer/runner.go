// Package installer provides software installation backends.
package installer

type Result struct {
	Output string
	Err    error
}

type Runner interface {
	Install(packages []string) Result
}
