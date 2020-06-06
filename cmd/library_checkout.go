package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			isCheckoutMostRecentLibrariesCommand, _ = libraryCheckoutCmd.Flags().GetBool("recent")
			isCheckoutProjectSpecifiedLibrariesCommand, _ = libraryCheckoutCmd.Flags().GetBool("project")
		},
	)

	libraryCheckoutCmd.Flags().BoolVar(
		&isCheckoutMostRecentLibrariesCommand,
		"recent",
		false,
		"Converts all libraries to their most recent versions.",
	)

	libraryCheckoutCmd.Flags().BoolVar(
		&isCheckoutProjectSpecifiedLibrariesCommand,
		"project",
		false,
		"Converts all libraries to the versions specified in the current project's config file.",
	)

	libraryCmd.AddCommand(libraryCheckoutCmd)

}

var libraryCheckoutCmd = &cobra.Command{

	Use: "checkout",

	Short: "Checks out the latest version of all libraries, or the versions specified for a particular project.",

	Long: "Checks out the latest version of all libraries, or the versions specified for a particular project.",

	Args: cobra.OnlyValidArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if isCheckoutMostRecentLibrariesCommand || isCheckoutProjectSpecifiedLibrariesCommand {
			var err error
			if isCheckoutMostRecentLibrariesCommand {
				err = checkoutMostRecentLibraries()
			} else if isCheckoutProjectSpecifiedLibrariesCommand {
				err = checkoutProjectSpecifiedLibraries()
			}
			if err != nil {
				util.ExitWithError(err)
			}
		} else {
			util.Println(cmd.UsageString())
		}
	},
}

var isCheckoutMostRecentLibrariesCommand bool
var isCheckoutProjectSpecifiedLibrariesCommand bool

func checkoutMostRecentLibraries() (err error) {

	globalConfig, err := configManager.GetGlobalConfig()
	if err != nil {
		return
	}
	libraryConfigList := globalConfig.Libraries

	for _, libraryConfig := range libraryConfigList {

		gitManager := util.NewGitManager(libraryConfig.FilePath)

		err = gitManager.Checkout("master")
		if err != nil {
			return
		}

		libraryConfig.CurrentGitCommitId = libraryConfig.MostRecentGitCommitId
		err = configManager.SaveGlobalConfig()
		if err != nil {
			return
		}

	}

	return

}

func checkoutProjectSpecifiedLibraries() (err error) {

	projectConfig, globalConfig, err := configManager.GetProjectAndGlobalConfigs()
	if err != nil {
		return
	}
	libraryProjectConfigList := projectConfig.Libraries
	libraryGlobalConfigList := globalConfig.Libraries

	for _, libraryProjectConfig := range libraryProjectConfigList {

		for _, libraryGlobalConfig := range libraryGlobalConfigList {

			if libraryProjectConfig.FilePath == libraryGlobalConfig.FilePath {

				libraryProjectConfig.MostRecentGitCommitId = libraryGlobalConfig.MostRecentGitCommitId
				libraryGlobalConfig.CurrentGitCommitId = libraryProjectConfig.CurrentGitCommitId

				gitManager := util.NewGitManager(libraryProjectConfig.FilePath)

				err = gitManager.Checkout(libraryProjectConfig.CurrentGitCommitId)
				if err != nil {
					return
				}

				err = configManager.SaveGlobalConfig()
				if err != nil {
					return
				}

				err = configManager.SaveProjectConfig()
				if err != nil {
					return
				}

			}

		}

	}

	return

}
