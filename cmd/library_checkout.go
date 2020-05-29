package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
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
			cmd.Help()
		}
	},
}

var isCheckoutMostRecentLibrariesCommand bool
var isCheckoutProjectSpecifiedLibrariesCommand bool

func checkoutMostRecentLibraries() (err error) {

	libraryConfigList := config.MppmGlobalConfig.Libraries

	for _, libraryConfig := range libraryConfigList {

		err = util.ExecuteGitCommandInDirectory(libraryConfig.FilePath, "checkout", libraryConfig.MostRecentGitCommitId)
		if err != nil {
			return
		}

		libraryConfig.CurrentGitCommitId = libraryConfig.MostRecentGitCommitId
		err = config.MppmGlobalConfig.SaveAsGlobalConfig()
		if err != nil {
			return
		}

	}

	return

}

func checkoutProjectSpecifiedLibraries() (err error) {

	config.LoadMppmProjectConfig()
	libraryProjectConfigList := config.MppmProjectConfig.Libraries
	libraryGlobalConfigList := config.MppmGlobalConfig.Libraries

	for _, libraryProjectConfig := range libraryProjectConfigList {

		for _, libraryGlobalConfig := range libraryGlobalConfigList {

			if libraryProjectConfig.FilePath == libraryGlobalConfig.FilePath {

				libraryProjectConfig.MostRecentGitCommitId = libraryGlobalConfig.MostRecentGitCommitId
				libraryGlobalConfig.CurrentGitCommitId = libraryProjectConfig.CurrentGitCommitId

				err = util.ExecuteGitCommandInDirectory(libraryProjectConfig.FilePath, "checkout", libraryProjectConfig.CurrentGitCommitId)
				if err != nil {
					return
				}

				err = config.MppmGlobalConfig.SaveAsGlobalConfig()
				if err != nil {
					return
				}

				err = config.MppmProjectConfig.SaveAsProjectConfig()
				if err != nil {
					return
				}

			}

		}

	}

	return

}
