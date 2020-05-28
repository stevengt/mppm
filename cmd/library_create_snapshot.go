package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	libraryCmd.AddCommand(libraryCreateSnapshotCmd)

}

var libraryCreateSnapshotCmd = &cobra.Command{

	Use: "create-snapshot",

	Short: "Creates a snapshot of all libraries (folders) currently tracked globally on your system.",

	Long: "Creates a snapshot of all libraries (folders) currently tracked globally on your system.",

	Args: cobra.NoArgs,
}
