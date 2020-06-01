package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func showSupportedFileTypes() {
	for _, filePatternsConfig := range config.GetFilePatternsConfigList() {
		filePatternsConfig.Print()
	}
}
