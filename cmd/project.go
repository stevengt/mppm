package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	cobra.OnInitialize(
		func() {
			isPreviewCommand, _ = projectCmd.PersistentFlags().GetBool("preview")
		},
	)

	projectCmd.PersistentFlags().BoolVar(
		&isPreviewCommand,
		"preview",
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
