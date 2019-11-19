package components

type ComponentInstaller interface {
	install(installDir string)
}

type ComponentInfo struct {
	Type            ComponentType
	DownloadURL     string
	InstallationDir string
}
