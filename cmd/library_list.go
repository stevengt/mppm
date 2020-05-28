package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	libraryCmd.AddCommand(libraryListCmd)

}

var libraryListCmd = &cobra.Command{

	Use: "list",

	Short: "Lists all libraries (folders) currently tracked globally on your system.",

	Long: "Lists all libraries (folders) currently tracked globally on your system.",

	Args: cobra.NoArgs,
}
