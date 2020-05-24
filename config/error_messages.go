package config

import (
	"encoding/json"
	"fmt"
)

var openingMppmProjectConfigFileErrorMessage = "There was a problem while opening the mppm config file.\n" +
	"If the file doesn't exist, try running 'mppm project init' first.\n"

func getInvalidMppmProjectConfigFileErrorMessage(jsonUnmarshalError error) (errorMessage string) {

	errorMessageTemplate := "The mppm config file %s is invalid." +
		"\n\t%s\n\n" +
		"An example valid config file is formatted like this:" +
		"\n\n%s\n"

	defaultMppmProjectConfigAsJson, err := json.Marshal(GetDefaultMppmProjectConfig())
	if err != nil {
		jsonMarshalErrorMessage := "Something went wrong while loading the default Mppm project config: " + err.Error()
		defaultMppmProjectConfigAsJson = []byte(jsonMarshalErrorMessage)
	}

	errorMessage = fmt.Sprintf(
		errorMessageTemplate,
		MppmProjectConfigFileName,
		jsonUnmarshalError.Error(),
		string(defaultMppmProjectConfigAsJson),
	)

	return

}
