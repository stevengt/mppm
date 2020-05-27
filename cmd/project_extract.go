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

	Use: "extract",

	Short: "Extracts all binary files of supported types into plain-text files, such as XML.",

	Long: `Extracts all binary files of supported types into plain-text files, such as XML.
This saves space within the git repository, and enables easier side-by-side comparison of different versions of the files.
			
Note that the original files are not stored in git directly.
To restore the original files, run 'mppm project restore'.`,

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if err := extractAllCompressedFiles(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func extractAllCompressedFiles() (err error) {

	filePatternsConfigList := make([]*config.FilePatternsConfig, 0)

	projectApplicationConfigs := config.MppmProjectConfig.Applications
	for _, projectApplicationConfig := range projectApplicationConfigs {
		applicationName := projectApplicationConfig.Name
		applicationVersion := projectApplicationConfig.Version
		for _, supportedApplication := range config.SupportedApplications {
			if supportedApplication.Name == applicationName {
				filePatternsConfig := supportedApplication.FilePatternConfigs[applicationVersion]
				filePatternsConfigList = append(filePatternsConfigList, filePatternsConfig)
			}
		}
	}

	for _, filePatternsConfig := range filePatternsConfigList {
		err = extractAllGzippedXmlFiles(filePatternsConfig)
		if err != nil {
			return
		}
	}

	return
}

func extractAllGzippedXmlFiles(filePatternsConfig *config.FilePatternsConfig) (err error) {
	gzippedXmlFileExtensions := filePatternsConfig.GzippedXmlFileExtensions
	for _, fileExtension := range gzippedXmlFileExtensions {
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

	for _, originalFileName := range fileNames {
		gzippedFileName := originalFileName + ".xml.gz"
		newFileName := originalFileName + ".xml"

		if isPreviewCommand {
			printExtractPreviewMessage(originalFileName, newFileName)
		} else {

			err = util.CopyFile(originalFileName, gzippedFileName)
			if err != nil {
				return
			}

			err = util.GunzipFile(gzippedFileName)
			if err != nil {
				return
			}

		}

	}

	return

}

func printExtractPreviewMessage(originalFileName string, newFileName string) {
	fmt.Println(originalFileName + " will be extracted to " + newFileName)
}
