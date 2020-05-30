package util

var gitManagerFactory GitManagerCreator = &GitShellCommandManagerCreator{}

func NewGitManager(repoFilePath string) GitManager {
	return gitManagerFactory.NewGitManager(repoFilePath)
}

type GitManagerCreator interface {
	NewGitManager(repoFilePath string) GitManager
}

type GitShellCommandManagerCreator struct{}

func (gitShellCommandManagerCreator *GitShellCommandManagerCreator) NewGitManager(repoFilePath string) GitManager {
	return &GitShellCommandManager{
		RepositoryDirectoryPath: repoFilePath,
	}
}

type GitManager interface {
	Init() (err error)
	Add(args ...string) (err error)
	Commit(args ...string) (err error)
	Checkout(args ...string) (err error)
	RevParse(args ...string) (stdout string, err error)
	LfsInstall() (err error)
	LfsTrack(args ...string) (err error)
	AddAllAndCommit(commitMessage string) (err error)
}

type GitShellCommandManager struct {
	RepositoryDirectoryPath string
}

func (gitShellCommandManager *GitShellCommandManager) executeGitShellCommand(gitCommandName string, args ...string) (err error) {
	commandName := "git"
	commandArgs := append(
		[]string{
			"-C",
			gitShellCommandManager.RepositoryDirectoryPath,
			gitCommandName,
		},
		args...,
	)
	err = ExecuteShellCommand(commandName, commandArgs...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) Init() (err error) {
	err = gitShellCommandManager.executeGitShellCommand("init")
	return
}

func (gitShellCommandManager *GitShellCommandManager) Add(args ...string) (err error) {
	err = gitShellCommandManager.executeGitShellCommand("add", args...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) Commit(args ...string) (err error) {
	err = gitShellCommandManager.executeGitShellCommand("commit", args...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) Checkout(args ...string) (err error) {
	err = gitShellCommandManager.executeGitShellCommand("checkout", args...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) RevParse(args ...string) (stdout string, err error) {
	commandName := "git"
	commandArgs := append(
		[]string{
			"-C",
			gitShellCommandManager.RepositoryDirectoryPath,
			"rev-parse",
		},
		args...,
	)
	stdout, err = ExecuteShellCommandAndReturnOutput(commandName, commandArgs...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) LfsInstall() (err error) {
	err = gitShellCommandManager.executeGitShellCommand("lfs", "install")
	return
}

func (gitShellCommandManager *GitShellCommandManager) LfsTrack(args ...string) (err error) {
	gitCommandName := "lfs"
	gitCommandArgs := append(
		[]string{
			"track",
		},
		args...,
	)
	err = gitShellCommandManager.executeGitShellCommand(gitCommandName, gitCommandArgs...)
	return
}

func (gitShellCommandManager *GitShellCommandManager) AddAllAndCommit(commitMessage string) (err error) {

	err = gitShellCommandManager.Add("-A", ".")
	if err != nil {
		return
	}

	err = gitShellCommandManager.Commit("-m", commitMessage)
	if err != nil {
		return
	}

	return

}
