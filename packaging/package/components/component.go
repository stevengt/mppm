package components

import "github.com/stevengt/mppm/packaging/package/components/installation/steps"

// Download URL's and installation directories should be specified
// in the InstallSteps. There should also (eventually) be a global
// default configuration for where to store certain types of components.

// Components should not have dependencies. Rather, dependencies should
// be represented in packages.

// Any components which can be downloaded/installed individually should
// have install steps like this:
//		- ComponentInfo install steps:
//			- Download from URL
//			- Extract files
//			- Move files
//			- Delete auxilliary files
// Any collections of components (e.g., a collection of samples)
// which share a single download URL should have install steps like this:
//		- ComponentCollectionInfo install steps:
//			- Download from URL
//			- Extract files
//			- For each component in collection:
//				- Run component's install steps
//			- Delete auxilliary files

type ComponentInfo struct {
	Version        string
	Description    string
	InstallSteps   []steps.ComponentInstallStepRunner
	UninstallSteps []steps.ComponentInstallStepRunner
}

type ComponentCollectionInfo struct {
	ComponentInfo
	Components []ComponentInfo
}
