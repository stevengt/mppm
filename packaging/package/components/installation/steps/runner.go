package steps

import (
	"github.com/stevengt/mppm/packaging/package/components/installation/steps/extract"
	"github.com/stevengt/mppm/packaging/package/components/installation/steps/mount"
	"github.com/stevengt/mppm/packaging/package/components/installation/steps/run"
	"github.com/stevengt/mppm/packaging/package/components/installation/steps/unmount"
)

type ComponentInstallStepRunner interface {
	Run() (err error)
}

func GetComponentInstallStepRunner(stepTypeName string, args ...string) (runner ComponentInstallStepRunner, err error) {
	stepType := GetStepTypeFromString(stepTypeName)
	switch stepType {
	case CopyFile:
		runner = CopyFileStepRunner{}
	case DeleteFile:
		runner = DeleteFileStepRunner{}
	case DownloadFile:
		runner = DownloadFileStepRunner{}
	case ExtractTarFile:
		runner = extract.ExtractTarFileStepRunner{}
	case ExtractZipFile:
		runner = extract.ExtractZipFileStepRunner{}
	case MountOsxDmgFile:
		runner = mount.MountOsxDmgFileStepRunner{}
	case MoveFile:
		runner = MoveFileStepRunner{}
	case RenameFile:
		runner = RenameFileStepRunner{}
	case RunExecutableFile:
		runner = run.RunExecutableFileStepRunner{}
	case RunOsxPkgInstallFile:
		runner = run.RunOsxPkgInstallFileStepRunner{}
	case UnmountOsxDmgFile:
		runner = unmount.UnmountOsxDmgFileStepRunner{}
	}
	return
}
