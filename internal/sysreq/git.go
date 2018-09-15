package sysreq

import (
	"os/exec"
)

type Git struct {
	cmd *exec.Cmd
}

func (g Git) Name() string {
	return "Git"
}

func (g Git) Has() bool {
	err := g.cmd.Run()

	return err == nil
}

func (g Git) InstallationDescription() string {
	return "Download from: https://git-scm.com/ for Mac OSX `brew install git`"
}

func CreateGitReq() Git {
	return Git{
		cmd: exec.Command("git", "--version"),
	}
}
