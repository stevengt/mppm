package applications

import (
	"fmt"
	"sort"
	"strings"
)

// Returns a list of *FilePatternsConfig, including all non-application-specific configs
// and all supported application-specific configs.
func GetFilePatternsConfigList() (filePatternsConfigList []*FilePatternsConfig) {
	filePatternsConfigList = append(
		GetNonApplicationSpecificFilePatternsConfigList(),
		GetApplicationSpecificFilePatternsConfigList()...,
	)
	return
}

// Returns a single *FilePatternsConfig containing the aggregate of all
// non-application-specific configs and all supported application-specific configs.
func GetAllFilePatternsConfig() (allFilePatternsConfig *FilePatternsConfig) {
	allFilePatternsConfig = NewFilePatternsConfig()
	for _, filePatternsConfig := range GetFilePatternsConfigList() {
		allFilePatternsConfig = allFilePatternsConfig.AppendAll(filePatternsConfig)
	}
	return
}

// ------------------------------------------------------------------------------

type FilePatternsConfig struct {
	Name                     string
	GitIgnorePatterns        []string
	GitLfsTrackPatterns      []string
	GzippedXmlFileExtensions []string // List of file extensions that represent Gzipped XML files.
}

func NewFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {
	return &FilePatternsConfig{
		Name:                     "",
		GitIgnorePatterns:        make([]string, 0),
		GitLfsTrackPatterns:      make([]string, 0),
		GzippedXmlFileExtensions: make([]string, 0),
	}
}

func (config *FilePatternsConfig) Print() {
	fmt.Print(config.Name + "\n\n")
	fmt.Print("\tGit Ignore Patterns \n\t\t")
	fmt.Println(strings.Join(config.GitIgnorePatterns, "\n\t\t"))
	fmt.Print("\tGit LFS Track Patterns \n\t\t")
	fmt.Println(strings.Join(config.GitLfsTrackPatterns, "\n\t\t"))
	fmt.Print("\tGzipped XML File Types \n\t\t")
	fmt.Println(strings.Join(config.GzippedXmlFileExtensions, "\n\t\t"))
}

func (config *FilePatternsConfig) SortAllLists() {
	sort.Strings(config.GitIgnorePatterns)
	sort.Strings(config.GitLfsTrackPatterns)
	sort.Strings(config.GzippedXmlFileExtensions)
}

func (config1 *FilePatternsConfig) AppendAll(config2 *FilePatternsConfig) (filePatternsConfig *FilePatternsConfig) {
	config1.GitIgnorePatterns = appendUnique(config1.GitIgnorePatterns, config2.GitIgnorePatterns)
	config1.GitLfsTrackPatterns = appendUnique(config1.GitLfsTrackPatterns, config2.GitLfsTrackPatterns)
	config1.GzippedXmlFileExtensions = appendUnique(config1.GzippedXmlFileExtensions, config2.GzippedXmlFileExtensions)
	filePatternsConfig = config1
	return
}

// ------------------------------------------------------------------------------

func appendUnique(list1 []string, list2 []string) (newList []string) {
	newList = make([]string, 0)
	uniqueVals := make(map[string]bool)
	for _, list1Val := range list1 {
		uniqueVals[list1Val] = true
	}
	for _, list2Val := range list2 {
		uniqueVals[list2Val] = true
	}
	for val, _ := range uniqueVals {
		newList = append(newList, val)
	}
	return
}
