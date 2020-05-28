package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/stevengt/mppm/config"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/util"
)

func init() {
	projectCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{

	Use: "restore",

	Short: "Restores all plain-text files of supported types to their original binary files.",

	Long: `Restores all plain-text files of supported types to their original binary files.
			
Note that the original files are not stored in git directly.
To extract them into plain-text files for use in git, run 'mppm project extract'.`,

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if err := restoreAllUncompressedFilesToOriginalCompressedFiles(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func restoreAllUncompressedFilesToOriginalCompressedFiles() (err error) {

	filePatternsConfig := config.GetAllFilePatternsConfigFromProjectConfig()

	err = restoreAllGzippedXmlFiles(filePatternsConfig)
	if err != nil {
		return
	}

	return
}

func restoreAllGzippedXmlFiles(filePatternsConfig *config.FilePatternsConfig) (err error) {
	gzippedXmlFileExtensions := filePatternsConfig.GzippedXmlFileExtensions
	for _, fileExtension := range gzippedXmlFileExtensions {
		err = restoreAllGzippedXmlFilesWithExtension(fileExtension)
		if err != nil {
			return
		}
	}
	return
}

func restoreAllGzippedXmlFilesWithExtension(fileExtension string) (err error) {

	fileNames, err := util.GetAllFileNamesWithExtension(fileExtension + ".xml")
	if err != nil {
		return
	}

	for _, originalFileName := range fileNames {

		newFileName := strings.TrimSuffix(originalFileName, ".xml")

		if isPreviewCommand {
			printRestorePreviewMessage(originalFileName, newFileName)
		} else {

			err = util.CopyFile(originalFileName, newFileName)
			if err != nil {
				return
			}

			err = util.GzipFile(newFileName)
			if err != nil {
				return
			}

			err = os.Rename(newFileName+".gz", newFileName)
			if err != nil {
				return
			}

		}

	}

	return

}

func printRestorePreviewMessage(originalFileName string, newFileName string) {
	fmt.Println(newFileName + " will be restored from " + originalFileName)
}
