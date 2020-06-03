package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			configManager = config.MppmConfigFileManager
			isShowSupportedFileTypesCommand, _ = rootCmd.Flags().GetBool("show-supported")
		},
	)

	rootCmd.Flags().BoolVarP(
		&isShowSupportedFileTypesCommand,
		"show-supported",
		"s",
		false,
		"Shows what file types are supported by mppm.",
	)

}

var rootCmd = &cobra.Command{

	Version: config.Version,

	Use: "mppm",

	Short: "Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects.",

	Long: `Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects, such as:

	- Simplified version control using 'git' and 'git-lfs'.
	- Extraction of 'Ableton Live Set' files to/from raw XML files.`,

	Args: cobra.OnlyValidArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if isShowSupportedFileTypesCommand {
			showSupportedFileTypes()
		} else {
			cmd.Help()
		}
	},
}

var configManager config.MppmConfigManager
var isShowSupportedFileTypesCommand bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.ExitWithError(err)
	}
}

func showSupportedFileTypes() {
	for _, filePatternsConfig := range applications.GetFilePatternsConfigList() {
		filePatternsConfig.Print()
	}
}
