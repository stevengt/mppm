package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use: "mppm",

	Short: "Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects.",

	Long: `Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects, such as:

	- Simplified version control using 'git' and 'git-lfs'.
	- Extraction of 'Ableton Live Set' files to/from raw XML files.`,

	Args: cobra.MinimumNArgs(1),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
