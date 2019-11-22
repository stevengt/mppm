package steps

import "errors"

type DownloadFileStepRunner struct{}

func (stepRunner DownloadFileStepRunner) Run() (err error) {
	err = errors.New("Not Implemented.")
	return
}
