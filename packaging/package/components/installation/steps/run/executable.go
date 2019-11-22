package run

import "errors"

type RunExecutableFileStepRunner struct{}

func (stepRunner RunExecutableFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
