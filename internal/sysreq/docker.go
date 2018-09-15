package sysreq

import (
	"os/exec"
)

type Docker struct {
	cmd *exec.Cmd
}

func (d Docker) Name() string {
	return "Docker"
}

func (d Docker) Has() bool {
	err := d.cmd.Run()

	return err == nil
}

func (d Docker) InstallationDescription() string {
	return "Download from: https://git-scm.com/ for Mac OSX `brew install docker`"
}

func CreateDockerReq() Docker {
	return Docker{
		cmd: exec.Command("docker", "--version"),
	}
}
