package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
)

func init() {

	existingMppmProjectConfig := config.MppmProjectConfig
	if existingMppmProjectConfig != nil {
		isCompatible, installedVersion, configVersion := existingMppmProjectConfig.IsCompatibleWithInstalledMppmVersion()
		if !isCompatible {
			fmt.Println("ERROR: Installed mppm version " + installedVersion +
				" is not compatible with this project's configured version " + configVersion)
			os.Exit(1)
		}
	}

	cobra.OnInitialize(
		func() {
			isPreviewCommand, _ = projectCmd.PersistentFlags().GetBool("preview")
		},
	)

	projectCmd.PersistentFlags().BoolVarP(
		&isPreviewCommand,
		"preview",
		"p",
		false,
		"Shows what files will be affected without actually making changes.",
	)

	rootCmd.AddCommand(projectCmd)

}

var isPreviewCommand bool

var projectCmd = &cobra.Command{

	Use: "project",

	Short: "Provides utilities for managing a specific project.",

	Long: "Provides utilities for managing a specific project.",

	Args: cobra.MinimumNArgs(1),
}
