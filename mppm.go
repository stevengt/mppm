package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/stevengt/mppm/versioning"
	"github.com/stevengt/mppm/versioning/project"
)

const HelpMessage string = "Please run using 'mppm init <project type>' or 'mppm git <git args>'."

func main() {

	if len(os.Args) < 3 {
		fmt.Println(HelpMessage)
		os.Exit(1)
	}

	switch os.Args[1] {

	case "init":

		var err error
		projectType := os.Args[2]

		switch projectType {
		case string(project.Ableton):
			versioner := &versioning.AbletonVersioner{}
			err = versioner.Init()
		default:
			err = errors.New("Please specify a valid project type (e.g., 'ableton').")
		}

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	case "git":

		gitArgs := os.Args[2:]
		config, err := project.LoadMppmConfig()

		if err == nil {
			switch config.ProjectType {
			case project.Ableton:
				versioner := &versioning.AbletonVersioner{}
				err = versioner.Git(gitArgs...)
			default:
				err = errors.New("A valid project type was not found in the config file '" + project.ConfigFileName + "'.")
			}
		}

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	default:
		fmt.Println(HelpMessage)
		os.Exit(1)
	}
}
