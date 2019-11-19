package components

type ComponentType int32

const (
	DAWSpecific ComponentType = iota
	Driver      ComponentType = iota
	Patch       ComponentType = iota
	Plugin      ComponentType = iota
	Preset      ComponentType = iota
	Sample      ComponentType = iota
	Soundbank   ComponentType = iota
)
