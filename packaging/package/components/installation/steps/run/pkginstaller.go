package run

import "errors"

type RunOsxPkgInstallFileStepRunner struct{}

func (stepRunner RunOsxPkgInstallFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
