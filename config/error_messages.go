package config

import (
	"encoding/json"
	"fmt"
)

func getOpeningMppmProjectConfigFileErrorMessage(err error) (errorMessage string) {
	errorMessageTemplate := `
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
%s
`
	errorMessage = fmt.Sprintf(errorMessageTemplate, err.Error())
	return
}

func getInvalidMppmProjectConfigFileErrorMessage(jsonUnmarshalError error) (errorMessage string) {

	errorMessageTemplate := `
The mppm config file %s is invalid.
	%s

An example valid config file is formatted like this:

%s
`

	defaultMppmProjectConfigAsJson, err := json.Marshal(GetDefaultMppmProjectConfig())
	if err != nil {
		jsonMarshalErrorMessage := "Something went wrong while loading the default Mppm project config: " + err.Error()
		defaultMppmProjectConfigAsJson = []byte(jsonMarshalErrorMessage)
	}

	errorMessage = fmt.Sprintf(
		errorMessageTemplate,
		MppmConfigFileName,
		jsonUnmarshalError.Error(),
		string(defaultMppmProjectConfigAsJson),
	)

	return

}

func getIncompatibleMppmVersionErrorMessage(installedVersion string, configVersion string) (errorMessage string) {
	errorMessage = fmt.Sprintf(
		"Installed mppm version %s is not compatible with this project's configured version %s",
		installedVersion,
		configVersion,
	)
	return
}

func getUnsupportedApplicationErrorMessage(application *ApplicationConfig) (errorMessage string) {
	errorMessageTemplate := `
Found unsupported application %s %s in config file %s
To see what applications are supported, please run 'mppm --show-supported'.
`
	errorMessage = fmt.Sprintf(
		errorMessageTemplate,
		application.Name,
		application.Version,
		MppmConfigFileName,
	)
	return
}
