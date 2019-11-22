package unmount

import "errors"

type UnmountOsxDmgFileStepRunner struct{}

func (stepRunner UnmountOsxDmgFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
