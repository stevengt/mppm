package extract

import "errors"

type ExtractTarFileStepRunner struct{}

func (stepRunner ExtractTarFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
