package ableton

type AbletonComponentType int32

const (
	AnalysisFile     AbletonComponentType = iota
	DevicePreset     AbletonComponentType = iota
	DeviceChain      AbletonComponentType = iota
	DeviceRack       AbletonComponentType = iota
	Groove           AbletonComponentType = iota
	LiveClip         AbletonComponentType = iota
	LivePack         AbletonComponentType = iota
	LiveProject      AbletonComponentType = iota
	LiveSet          AbletonComponentType = iota
	MaxForLiveDevice AbletonComponentType = iota
	MetaSound        AbletonComponentType = iota
	Skin             AbletonComponentType = iota
)
