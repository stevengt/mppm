package plugins

type PluginType int32

const (
	AudioUnit2 PluginType = iota
	VST2       PluginType = iota
	VST3       PluginType = iota
)
