package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/stevengt/mppm/versioning"
	"github.com/stevengt/mppm/versioning/project"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes version control settings for a project using git and git-lfs.",
	Long:  "Initializes version control settings for a project using git and git-lfs.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initProject(args[0])
	},
}

func initProject(projectType string) {
	var err error

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

}
