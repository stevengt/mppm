package extract

import "errors"

type ExtractZipFileStepRunner struct{}

func (stepRunner ExtractZipFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
