package configtest

import (
	"errors"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

var TestMppmConfigInfoAsJsonAndExpectedConfigFunctionResponses []*MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses = []*MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{

	// valid version, valid application name, valid application version
	&MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{
		ConfigAsJson:  []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"10"}]}`),
		ExpectedError: nil,
		ExpectedFilePatternsConfigList: []*applications.FilePatternsConfig{
			applications.AudioFilePatternsConfig,
			applications.Ableton10FilePatternsConfig,
		},
		ExpectedFilePatternsConfig: &applications.FilePatternsConfig{
			Name:              "",
			GitIgnorePatterns: []string{"Backup/", "*.als", "*.alc", "*.adv", "*.adg"},
			GitLfsTrackPatterns: []string{"*.flac", "*.iklax", "*.m4a", "*.alac", "*.au",
				"*.mpc", "*.ogg", "*.mogg", "*.tta", "*.wma", "*.aax", "*.act", "*.ivs",
				"*.aa", "*.dvf", "*.m4b", "*.nsf", "*.raw", "*.webm", "*.cda", "*.dct",
				"*.gsm", "*.dss", "*.msv", "*.nmf", "*.sln", "*.3gp", "*.aac", "*.voc",
				"*.wv", "*.m4p", "*.rm", "*.ape", "*.awb", "*.mmf", "*.oga", "*.opus",
				"*.rf64", "*.aiff", "*.amr", "*.vox", "*.wav", "*.8svx", "*.mp3", "*.ra",
				"*.ams", "*.amxd", "*.alp", "*.asd", "*.agr",
			},
			GzippedXmlFileExtensions: []string{"adv", "adg", "als", "alc"},
		},
	},

	// valid version, no applications
	&MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{
		ConfigAsJson:  []byte(`{"version":"1.0.0","applications":[]}`),
		ExpectedError: nil,
		ExpectedFilePatternsConfigList: []*applications.FilePatternsConfig{
			applications.AudioFilePatternsConfig,
		},
		ExpectedFilePatternsConfig: &applications.FilePatternsConfig{
			Name:              "",
			GitIgnorePatterns: []string{},
			GitLfsTrackPatterns: []string{"*.flac", "*.iklax", "*.m4a", "*.alac", "*.au",
				"*.mpc", "*.ogg", "*.mogg", "*.tta", "*.wma", "*.aax", "*.act", "*.ivs",
				"*.aa", "*.dvf", "*.m4b", "*.nsf", "*.raw", "*.webm", "*.cda", "*.dct",
				"*.gsm", "*.dss", "*.msv", "*.nmf", "*.sln", "*.3gp", "*.aac", "*.voc",
				"*.wv", "*.m4p", "*.rm", "*.ape", "*.awb", "*.mmf", "*.oga", "*.opus",
				"*.rf64", "*.aiff", "*.amr", "*.vox", "*.wav", "*.8svx", "*.mp3", "*.ra",
			},
			GzippedXmlFileExtensions: []string{},
		},
	},

	// invalid version, no applications
	&MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{
		ConfigAsJson:                   []byte(`{"version":"0.0.0","applications":[]}`),
		ExpectedError:                  errors.New("Installed mppm version 1.2.1 is not compatible with this project's configured version 0.0.0"),
		ExpectedFilePatternsConfigList: nil,
		ExpectedFilePatternsConfig:     nil,
	},

	// valid version, invalid application name
	&MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{
		ConfigAsJson: []byte(`{"version":"1.0.0","applications":[{"name":"Fake Application","version":"1"}]}`),
		ExpectedError: errors.New(`
Found unsupported application Fake Application 1 in config file .mppm.json
To see what applications are supported, please run 'mppm --show-supported'.
`),
		ExpectedFilePatternsConfigList: nil,
		ExpectedFilePatternsConfig:     nil,
	},

	// valid version, valid application name, invalid application version
	&MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses{
		ConfigAsJson: []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"1"}]}`),
		ExpectedError: errors.New(`
Found unsupported application Ableton 1 in config file .mppm.json
To see what applications are supported, please run 'mppm --show-supported'.
`),
		ExpectedFilePatternsConfigList: nil,
		ExpectedFilePatternsConfig:     nil,
	},
}

func InitMockFileSystemDelegaterWithDefaultConfigFiles() {

	_, projectConfigAsJson := GetTestMppmConfigInfo()
	_, globalConfigAsJson := GetTestMppmConfigInfo()
	InitMockFileSystemDelegaterWithConfigFiles(projectConfigAsJson, globalConfigAsJson)

}

func InitMockFileSystemDelegaterWithConfigFiles(projectConfigAsJson []byte, globalConfigAsJson []byte) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}

	projectConfigFilePath := config.MppmConfigFileName
	userHomeDirectoryFilePath, _ := mockFileSystemDelegater.UserHomeDir()
	globalConfigFilePath := mockFileSystemDelegater.JoinFilePath(userHomeDirectoryFilePath, config.MppmConfigFileName)

	projectConfigFile := utiltest.NewMockFile(projectConfigAsJson)
	globalConfigFile := utiltest.NewMockFile(globalConfigAsJson)

	mockFileSystemDelegater.Files = map[string]*utiltest.MockFile{
		projectConfigFilePath: projectConfigFile,
		globalConfigFilePath:  globalConfigFile,
	}

	util.FileSystemProxy = mockFileSystemDelegater
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()

}

func GetTestMppmConfigInfo() (testMppmConfigInfo *config.MppmConfigInfo, configAsJson []byte) {

	testMppmConfigInfo = &config.MppmConfigInfo{
		Version: "1.0.0",
		Applications: []*applications.ApplicationConfig{
			&applications.ApplicationConfig{
				Name:    "Ableton",
				Version: "10",
			},
		},
		Libraries: []*config.LibraryConfig{
			&config.LibraryConfig{
				FilePath:              "/home/testuser/library",
				MostRecentGitCommitId: "56789",
				CurrentGitCommitId:    "01234",
			},
		},
	}

	configAsJson = []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"10"}],"libraries":[{"location":"/home/testuser/library","most-recent-version":"56789","current-version":"01234"}]}`)

	return

}

// Contains the contents of a config file as a JSON byte-array, and the
// expected responses for functions that read from the config file.
type MppmConfigInfoAsJsonAndExpectedConfigFunctionResponses struct {
	ConfigAsJson                   []byte
	ExpectedError                  error
	ExpectedFilePatternsConfigList []*applications.FilePatternsConfig // Expected response from config.GetFilePatternsConfigListFromProjectConfig().
	ExpectedFilePatternsConfig     *applications.FilePatternsConfig   // Expected response from config.GetAllFilePatternsConfigFromProjectConfig().
}
