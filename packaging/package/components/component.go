package components

import "github.com/stevengt/mppm/packaging/package/components/installation/steps"

type ComponentInfo struct {
	DownloadURL     string
	InstallationDir string
	Version         string
	Description     string
	InstallSteps    steps.ComponentInstallStepRunner
	UninstallSteps  steps.ComponentInstallStepRunner
	Dependencies    []ComponentInfo
}
