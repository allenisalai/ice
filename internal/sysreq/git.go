package sysreq

import "os/exec"

type Git struct{}

func (g Git) Has() bool {
	cmd := exec.Command("git", "-v")
	err := cmd.Run()

	return err == nil
}

func (g Git) GetDescription() string {
	return "Git is a system requirement"
}