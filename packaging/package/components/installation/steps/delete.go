package steps

import "errors"

type DeleteFileStepRunner struct{}

func (stepRunner DeleteFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
