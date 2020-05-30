package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	libraryCmd.AddCommand(libraryRemoveCmd)

}

var libraryRemoveCmd = &cobra.Command{

	Use: "remove",

	Short: "Removes a library (folder) to track globally on your system.",

	Long: `Removes a library (folder) to track globally on your system.

All previous library changes will still be saved, but mppm will not track any future changes.
`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		err := removeLibrary(args[0])
		if err != nil {
			util.ExitWithError(err)
		}
	},
}

func removeLibrary(libraryFilePath string) (err error) {

	currentLibraries := make([]*config.LibraryConfig, 0)

	for _, libraryConfig := range configManager.GetGlobalConfig().Libraries {
		if libraryConfig.FilePath != libraryFilePath {
			currentLibraries = append(currentLibraries, libraryConfig)
		}
	}

	configManager.GetGlobalConfig().Libraries = currentLibraries
	err = configManager.SaveGlobalConfig()
	if err != nil {
		return
	}

	return

}
