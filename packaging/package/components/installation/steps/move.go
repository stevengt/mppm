package steps

import "errors"

type MoveFileStepRunner struct{}

func (stepRunner MoveFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
