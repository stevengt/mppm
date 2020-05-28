package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	libraryCmd.AddCommand(libraryRemoveCmd)

}

var libraryRemoveCmd = &cobra.Command{

	Use: "remove",

	Short: "Removes a library (folder) to track globally on your system.",

	Long: "Removes a library (folder) to track globally on your system.",

	Args: cobra.MinimumNArgs(1),
}
