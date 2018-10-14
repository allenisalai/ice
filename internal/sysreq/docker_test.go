package sysreq

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestDocker_Has(t *testing.T) {

	t.Run("Run good command", func(t *testing.T) {
		cs := []string{"-test.run=TestHelperProcess", "--"}
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

		d := Docker{
			cmd: cmd,
		}

		if d.Has() == false {
			t.Error("A valid command returned as not being successful")
		}
	})

	t.Run("Run Bad git command", func(t *testing.T) {
		cs := []string{"-test.run=TestBadHelperProcess", "--"}
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

		d := Docker{
			cmd: cmd,
		}

		if d.Has() == true {
			t.Error("An invalid command returned as being successful")
		}
	})
}

func TestCreateDockerReq(t *testing.T) {
	d := CreateDockerReq()

	if strings.HasSuffix(d.cmd.Path, "docker") == false {
		t.Errorf("Command should call docker: %s called instead", d.cmd.Path)
	}
}


func TestDocker_InstallationDescription(t *testing.T) {
	d := Docker{}

	if d.InstallationDescription() != "Download from: https://git-scm.com/ for Mac OSX `brew install docker`" {
		t.Errorf("Docker installation description is out of date: %s", d.InstallationDescription())
	}
}

func TestDocker_Name(t *testing.T) {
	d := Docker{}

	if d.Name() != "Docker" {
		t.Errorf("Docker name is incorrect %s", d.Name())
	}
}
