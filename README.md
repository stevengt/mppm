
# mppm
### Music Production Project Manager


## Requirements
- [git](https://git-scm.com)
- [git-lfs](https://git-lfs.github.com)

## Installation
`go get -u github.com/stevengt/mppm`

## Usage
```
$ mppm --help

Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects, such as:

	- Simplified version control using 'git' and 'git-lfs'.
	- Extraction of 'Ableton Live Set' files to/from raw XML files.

Usage:
  mppm [flags]
  mppm [command]

Available Commands:
  help        Help about any command
  library     Provides utilities for globally managing multiple libraries (folders).
  project     Provides utilities for managing a specific project.

Flags:
  -h, --help             help for mppm
  -s, --show-supported   Shows what file types are supported by mppm.
  -v, --version          version for mppm

Use "mppm [command] --help" for more information about a command.
```

```
$ mppm project --help

Provides utilities for managing a specific project.

Usage:
  mppm project [flags]
  mppm project [command]

Available Commands:
  extract     Extracts all binary files of supported types into plain-text files, such as XML.
  init        Initializes version control settings for a project using git and git-lfs.
  restore     Restores all plain-text files of supported types to their original binary files.

Flags:
  -c, --commit-all         Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.
  -h, --help               help for project
  -p, --preview            Shows what files will be affected without actually making changes.
  -u, --update-libraries   Updates the library versions in the project config file to match the
                           current versions in the global config file.
                           To see the global current versions, run 'mppm library --list'.

Use "mppm project [command] --help" for more information about a command.
```

```
$ mppm library --help

Provides utilities for globally managing multiple libraries (folders).

Specifically, you can specify a list of folders and periodically take "snapshots"
of their contents. mppm can then keep track of which "versions" of the folders
each of your projects depend on, at any given time.

This is useful in the case that you inadvertently change/delete a file that your project
depends on. While working on your project, you can simply revert to a previous snapshot
of your libraries to get the original file back. When you finish working on your project,
you can then restore your libraries to their most recent snapshot.

Libraries can be any collection of audio samples, plugins, presets, etc. that:
	- Your projects might "depend on".
	- You expect, in general, to update less frequently compared to projects.

Usage:
  mppm library [flags]
  mppm library [command]

Available Commands:
  add         Adds a library (folder) to track globally on your system.
  checkout    Checks out the latest version of all libraries, or the versions specified for a particular project.
  remove      Removes a library (folder) to track globally on your system.

Flags:
  -c, --commit-all   Commits (snapshots) all changes made to all libraries (folders) currently tracked globally on your system.
  -h, --help         help for library
  -l, --list         Lists all libraries (folders) currently tracked globally on your system.

Use "mppm library [command] --help" for more information about a command.
```
