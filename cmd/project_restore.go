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
	Use:   "restore",
	Short: "Restores all 'Ableton Live Set' XML files to their original '.als' files.",
	Long: `Restores all 'Ableton Live Set' XML files to their original '.als' files.
			
			Note that no '.als' files are stored in git directly. To extract them into XML files for use in git, run 'mppm project extract'.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := restoreAllUncompressedFilesToOriginalCompressedFiles(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func restoreAllUncompressedFilesToOriginalCompressedFiles() (err error) {

	err = restoreAllGzippedXmlFiles()
	if err != nil {
		return
	}

	return
}

func restoreAllGzippedXmlFiles() (err error) {
	gzippedXmlFileExtensions := config.GetAllFilePatternsConfig().GzippedXmlFileExtensions
	for i := 0; i < len(gzippedXmlFileExtensions); i++ {
		fileExtension := gzippedXmlFileExtensions[i]
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

	for i := 0; i < len(fileNames); i++ {

		originalFileName := fileNames[i]
		newFileName := strings.TrimSuffix(originalFileName, ".xml")

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

	return

}
