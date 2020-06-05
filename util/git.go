package util

var GitManagerFactory GitManagerCreator = &gitShellCommandProxyCreator{}

func NewGitManager(repoFilePath string) GitManager {
	return GitManagerFactory.NewGitManager(repoFilePath)
}

// ------------------------------------------------------------------------------

type GitManagerCreator interface {
	NewGitManager(repoFilePath string) GitManager
}

type gitShellCommandProxyCreator struct{}

func (proxyCreator *gitShellCommandProxyCreator) NewGitManager(repoFilePath string) GitManager {
	return &gitShellCommandProxy{
		RepositoryDirectoryPath: repoFilePath,
	}
}

// ------------------------------------------------------------------------------

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

type gitShellCommandProxy struct {
	RepositoryDirectoryPath string
}

func (proxy *gitShellCommandProxy) Init() (err error) {
	err = proxy.executeGitShellCommand("init")
	return
}

func (proxy *gitShellCommandProxy) Add(args ...string) (err error) {
	err = proxy.executeGitShellCommand("add", args...)
	return
}

func (proxy *gitShellCommandProxy) Commit(args ...string) (err error) {
	err = proxy.executeGitShellCommand("commit", args...)
	return
}

func (proxy *gitShellCommandProxy) Checkout(args ...string) (err error) {
	err = proxy.executeGitShellCommand("checkout", args...)
	return
}

func (proxy *gitShellCommandProxy) RevParse(args ...string) (stdout string, err error) {
	stdout, err = proxy.executeGitShellCommandAndReturnOutput("rev-parse", args...)
	return
}

func (proxy *gitShellCommandProxy) LfsInstall() (err error) {
	err = proxy.executeGitShellCommand("lfs", "install")
	return
}

func (proxy *gitShellCommandProxy) LfsTrack(args ...string) (err error) {
	gitCommandName := "lfs"
	gitCommandArgs := append(
		[]string{
			"track",
		},
		args...,
	)
	err = proxy.executeGitShellCommand(gitCommandName, gitCommandArgs...)
	return
}

func (proxy *gitShellCommandProxy) AddAllAndCommit(commitMessage string) (err error) {

	err = proxy.Add("-A", ".")
	if err != nil {
		return
	}

	err = proxy.Commit("-m", commitMessage)
	if err != nil {
		return
	}

	return

}

func (proxy *gitShellCommandProxy) executeGitShellCommand(gitCommandName string, args ...string) (err error) {
	stdout, err := proxy.executeGitShellCommandAndReturnOutput(gitCommandName, args...)
	if len(stdout) > 0 {
		Println(stdout)
	}
	return
}

func (proxy *gitShellCommandProxy) executeGitShellCommandAndReturnOutput(gitCommandName string, args ...string) (stdout string, err error) {
	commandName := "git"
	commandArgs := append(
		[]string{
			"-C",
			proxy.RepositoryDirectoryPath,
			gitCommandName,
		},
		args...,
	)
	stdout, err = ExecuteShellCommandAndReturnOutput(commandName, commandArgs...)
	return
}
