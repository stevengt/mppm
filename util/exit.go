package util

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	getExiter().ExitWithError(err)
}

func ExitWithErrorMessage(errorMessage string) {
	getExiter().ExitWithErrorMessage(errorMessage)
}

var exiterFactory ExiterCreator = &CurrentProcessExiterCreator{}
var exiter Exiter

func getExiter() Exiter {
	if exiter == nil {
		exiter = exiterFactory.NewExiter()
	}
	return exiter
}

type ExiterCreator interface {
	NewExiter() Exiter
}

type CurrentProcessExiterCreator struct{}

func (currentProcessExiterCreator *CurrentProcessExiterCreator) NewExiter() Exiter {
	return &CurrentProcessExiter{}
}

type Exiter interface {
	ExitWithError(err error)
	ExitWithErrorMessage(errorMessage string)
}

type CurrentProcessExiter struct{}

func (currentProcessExiter *CurrentProcessExiter) ExitWithError(err error) {
	currentProcessExiter.ExitWithErrorMessage(err.Error())
}

func (currentProcessExiter *CurrentProcessExiter) ExitWithErrorMessage(errorMessage string) {
	fmt.Println("ERROR: " + errorMessage)
	os.Exit(1)
}
