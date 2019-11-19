package plugins

import "github.com/stevengt/mppm/packaging/package/components"

type PluginComponent struct {
	components.ComponentBase
	Type PluginType
}
