package components

type ComponentType string

const (
	Collection  ComponentType = "Collection"
	DAWSpecific ComponentType = "DAWSpecific"
	Driver      ComponentType = "Driver"
	Patch       ComponentType = "Patch"
	Plugin      ComponentType = "Plugin"
	Preset      ComponentType = "Preset"
	Sample      ComponentType = "Sample"
	Soundbank   ComponentType = "Soundbank"
)

func GetComponentTypeFromString(typeName string) (componentType ComponentType) {
	switch typeName {
	case "Collection":
		componentType = Collection
	case "DAWSpecific":
		componentType = DAWSpecific
	case "Driver":
		componentType = Driver
	case "Patch":
		componentType = Patch
	case "Plugin":
		componentType = Plugin
	case "Preset":
		componentType = Preset
	case "Sample":
		componentType = Sample
	case "Soundbank":
		componentType = Soundbank
	}
	return
}
