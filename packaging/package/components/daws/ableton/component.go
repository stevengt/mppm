package ableton

import "github.com/stevengt/mppm/packaging/package/components"

type AbletonComponent struct {
	components.ComponentBase
	Type AbletonComponentType
}
