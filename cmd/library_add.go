package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	libraryCmd.AddCommand(libraryAddCmd)

}

var libraryAddCmd = &cobra.Command{

	Use: "add",

	Short: "Adds a library (folder) to track globally on your system.",

	Long: "Adds a library (folder) to track globally on your system.",

	Args: cobra.MinimumNArgs(1),
}
