package steps

import "errors"

type CopyFileStepRunner struct{}

func (stepRunner CopyFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
