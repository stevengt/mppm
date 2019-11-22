package steps

import "errors"

type RenameFileStepRunner struct{}

func (stepRunner RenameFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
