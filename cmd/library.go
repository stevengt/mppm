package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
)

func init() {

	rootCmd.AddCommand(libraryCmd)

}

var libraryCmd = &cobra.Command{

	Use: "library",

	Short: "Provides utilities for globally managing multiple libraries (folders).",

	Long: `Provides utilities for globally managing multiple libraries (folders).

Specifically, you can specify a list of folders and periodically take "snapshots"
of their contents. mppm can then keep track of which "versions" of the folders
each of your projects depend on, at any given time.

This is useful in the case that you inadvertently change/delete a file that your project
depends on. While working on your project, you can simply revert to a previous snapshot
of your libraries to get the original file back. When you finish working on your project,
you can then restore your libraries to their most recent snapshot.

Libraries can be any collection of audio samples, plugins, presets, etc. that:
	- Your projects might "depend on".
	- You expect, in general, to update less frequently compared to projects.
`,

	Args: cobra.MinimumNArgs(1),

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.LoadMppmGlobalConfig()
	},
}
