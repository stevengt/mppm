package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
)

func init() {

	libraryCmd.AddCommand(libraryListCmd)

}

var libraryListCmd = &cobra.Command{

	Use: "list",

	Short: "Lists all libraries (folders) currently tracked globally on your system.",

	Long: "Lists all libraries (folders) currently tracked globally on your system.",

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		listAllTrackedLibraries()
	},
}

func listAllTrackedLibraries() {
	for _, libraryConfig := range config.MppmGlobalConfig.Libraries {
		libraryConfig.Print()
	}
}
