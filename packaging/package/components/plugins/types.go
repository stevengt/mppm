package plugins

type PluginType string

const (
	AudioUnit2 PluginType = "AudioUnit2"
	VST2       PluginType = "VST2"
	VST3       PluginType = "VST3"
)

func GetPluginTypeFromString(typeName string) (pluginType PluginType) {
	switch typeName {
	case "AudioUnit2":
		pluginType = AudioUnit2
	case "VST2":
		pluginType = VST2
	case "VST3":
		pluginType = VST3
	}
	return
}
