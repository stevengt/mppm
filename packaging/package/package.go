package packaging

import "github.com/stevengt/mppm/packaging/package/components"

type PackageInfo struct {
	Components   []*components.ComponentBase
	Dependencies []*PackageInfo
}
