package util

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	ExitWithErrorMessage(err.Error())
}

func ExitWithErrorMessage(errorMessage string) {
	fmt.Println("ERROR: " + errorMessage)
	os.Exit(1)
}
