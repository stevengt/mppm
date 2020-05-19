package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mppm",
	Short: "Short for 'Music Production Project Manager', mppm enables easy version control of music production projects",
	Long:  "Short for 'Music Production Project Manager', mppm enables easy version control of music production projects",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
