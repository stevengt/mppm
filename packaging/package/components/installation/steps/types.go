package steps

type ComponentInstallStepType string

// These commands require specific file/folder names to be defined.
// Dynamically accessed files/folders are not supported in order
// to ensure reproducibility and safety.
const (
	CopyFile             ComponentInstallStepType = "CopyFile"
	DeleteFile           ComponentInstallStepType = "DeleteFile"
	DownloadFile         ComponentInstallStepType = "DownloadFile"
	ExtractTarFile       ComponentInstallStepType = "ExtractTarFile"
	ExtractZipFile       ComponentInstallStepType = "ExtractZipFile"
	MountOsxDmgFile      ComponentInstallStepType = "MountOsxDmgFile"
	MoveFile             ComponentInstallStepType = "MoveFile"
	RenameFile           ComponentInstallStepType = "RenameFile"
	RunExecutableFile    ComponentInstallStepType = "RunExecutableFile"
	RunOsxPkgInstallFile ComponentInstallStepType = "RunOsxPkgInstallFile"
	UnmountOsxDmgFile    ComponentInstallStepType = "UnmountOsxDmgFile"
)

func GetStepTypeFromString(stepTypeName string) (stepType ComponentInstallStepType) {
	switch stepTypeName {
	case "CopyFile":
		stepType = CopyFile
	case "DeleteFile":
		stepType = DeleteFile
	case "DownloadFile":
		stepType = DownloadFile
	case "ExtractTarFile":
		stepType = ExtractTarFile
	case "ExtractZipFile":
		stepType = ExtractZipFile
	case "MountOsxDmgFile":
		stepType = MountOsxDmgFile
	case "MoveFile":
		stepType = MoveFile
	case "RenameFile":
		stepType = RenameFile
	case "RunExecutableFile":
		stepType = RunExecutableFile
	case "RunOsxPkgInstallFile":
		stepType = RunOsxPkgInstallFile
	case "UnmountOsxDmgFile":
		stepType = UnmountOsxDmgFile
	}
	return
}
