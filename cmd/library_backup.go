package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	libraryCmd.AddCommand(libraryBackupCmd)

}

var libraryBackupCmd = &cobra.Command{

	Use: "backup",

	Short: "Creates a backup at the specified location of all libraries (folders) currently tracked globally on your system.",

	Long: "Creates a backup at the specified location of all libraries (folders) currently tracked globally on your system.",

	Args: cobra.MinimumNArgs(1),
}
