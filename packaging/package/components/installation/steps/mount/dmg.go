package mount

import "errors"

type MountOsxDmgFileStepRunner struct{}

func (stepRunner MountOsxDmgFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
