package util

import (
	"fmt"
	"os"
)

var CurrentProcessExiter Exiter = &currentProcessExiter{}

func ExitWithError(err error) {
	CurrentProcessExiter.ExitWithError(err)
}

func ExitWithErrorMessage(errorMessage string) {
	CurrentProcessExiter.ExitWithErrorMessage(errorMessage)
}

type Exiter interface {
	ExitWithError(err error)
	ExitWithErrorMessage(errorMessage string)
}

type currentProcessExiter struct{}

func (exiter *currentProcessExiter) ExitWithError(err error) {
	exiter.ExitWithErrorMessage(err.Error())
}

func (exiter *currentProcessExiter) ExitWithErrorMessage(errorMessage string) {
	fmt.Println("ERROR: " + errorMessage)
	os.Exit(1)
}
