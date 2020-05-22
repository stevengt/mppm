package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/util"
)

func init() {
	projectCmd.AddCommand(extractCmd)
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts all 'Ableton Live Set' files to raw XML files.",
	Long: `Extracts all 'Ableton Live Set' files to raw XML files.
			This saves space within the git repository, and enables easier side-by-side comparison of different versions of the files.
			
			Note that no '.als' files are stored in git directly. To restore the original files from XML, run 'mppm project restore'.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := copyAllAlsFilesToUncompressedXmlFiles(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func copyAllAlsFilesToUncompressedXmlFiles() (err error) {

	fileNames, err := util.GetAllFileNamesWithExtension("als")
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