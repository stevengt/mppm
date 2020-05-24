package config

import (
	"fmt"
	"strings"
)

type FilePatternsConfig struct {
	Name                     string
	GitIgnorePatterns        []string
	GitLfsTrackPatterns      []string
	GzippedXmlFileExtensions []string // List of file extensions that represent Gzipped XML files.
}

var FilePatternsConfigList = []*FilePatternsConfig{
	AudioFilePatternsConfig,
	Ableton10FilePatternsConfig,
}

func GetAllFilePatternsConfig() (allFilePatternsConfig *FilePatternsConfig) {
	allFilePatternsConfig = newFilePatternsConfig()
	for _, filePatternsConfig := range FilePatternsConfigList {
		allFilePatternsConfig = allFilePatternsConfig.appendAll(filePatternsConfig)
	}
	return
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

func newFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {
	return &FilePatternsConfig{
		Name:                     "",
		GitIgnorePatterns:        make([]string, 0),
		GitLfsTrackPatterns:      make([]string, 0),
		GzippedXmlFileExtensions: make([]string, 0),
	}
}

func (config1 *FilePatternsConfig) appendAll(config2 *FilePatternsConfig) (filePatternsConfig *FilePatternsConfig) {
	config1.GitIgnorePatterns = appendUnique(config1.GitIgnorePatterns, config2.GitIgnorePatterns)
	config1.GitLfsTrackPatterns = appendUnique(config1.GitLfsTrackPatterns, config2.GitLfsTrackPatterns)
	config1.GzippedXmlFileExtensions = appendUnique(config1.GzippedXmlFileExtensions, config2.GzippedXmlFileExtensions)
	filePatternsConfig = config1
	return
}

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
