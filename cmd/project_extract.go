package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {
	projectCmd.AddCommand(extractCmd)
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts all binary files of supported types into plain-text files, such as XML.",
	Long: `Extracts all binary files of supported types into plain-text files, such as XML.
			This saves space within the git repository, and enables easier side-by-side comparison of different versions of the files.
			
			Note that the original files are not stored in git directly. To restore the original files, run 'mppm project restore'.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := extractAllCompressedFiles(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func extractAllCompressedFiles() (err error) {

	err = extractAllGzippedXmlFiles()
	if err != nil {
		return
	}

	return
}

func extractAllGzippedXmlFiles() (err error) {
	gzippedXmlFileExtensions := config.GetAllFilePatternsConfig().GzippedXmlFileExtensions
	for i := 0; i < len(gzippedXmlFileExtensions); i++ {
		fileExtension := gzippedXmlFileExtensions[i]
		err = extractAllGzippedXmlFilesWithExtension(fileExtension)
		if err != nil {
			return
		}
	}
	return
}

func extractAllGzippedXmlFilesWithExtension(fileExtension string) (err error) {

	fileNames, err := util.GetAllFileNamesWithExtension(fileExtension)
	if err != nil {
		return
	}

	for i := 0; i < len(fileNames); i++ {
		originalFileName := fileNames[i]
		newFileName := fileNames[i] + ".xml.gz"

		err = util.CopyFile(originalFileName, newFileName)
		if err != nil {
			return
		}

		err = util.GunzipFile(newFileName)
		if err != nil {
			return
		}
	}

	return

}
