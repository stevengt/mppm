package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Provides utilities for managing a specific project.",
	Long:  "Provides utilities for managing a specific project.",
	Args:  cobra.MinimumNArgs(1),
}
