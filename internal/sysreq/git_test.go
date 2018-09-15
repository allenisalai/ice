package sysreq

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGit_Has(t *testing.T) {

	t.Run("Run good command", func(t *testing.T) {
		cs := []string{"-test.run=TestHelperProcess", "--"}
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

		g := Git{
			cmd: cmd,
		}

		if g.Has() == false {
			t.Error("A valid command returned as not being successful")
		}
	})

	t.Run("Run Bad git command", func(t *testing.T) {
		cs := []string{"-test.run=TestBadHelperProcess", "--"}
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

		g := Git{
			cmd: cmd,
		}

		if g.Has() == true {
			t.Error("An invalid command returned as being successful")
		}
	})
}

func TestCreate(t *testing.T) {
	g := CreateGitReq()

	if strings.HasSuffix(g.cmd.Path, "git") == false {
		t.Errorf("Command should call git: %s called instead", g.cmd.Path)
	}
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		return
	}
}

func TestBadHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		os.Exit(1)
	}
}

func TestGit_InstallationDescription(t *testing.T) {
	g := Git{}

	if g.InstallationDescription() != "Download from: https://git-scm.com/ for Mac OSX `brew install git`" {
		t.Errorf("Git installation description is out of date: %s", g.InstallationDescription())
	}
}

func TestGit_Name(t *testing.T) {
	g := Git{}

	if g.Name() != "Git" {
		t.Errorf("Git name is incorrect %s", g.Name())
	}
}
