package ice

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

type InstalledService struct {
	Name string
	Dir  string
}

func GetInstalledServices(codeDir string) []InstalledService {
	var services []InstalledService

	folders, _ := ioutil.ReadDir(codeDir)
	for _, f := range folders {
		if f.IsDir() {
			services = append(services, InstalledService{
				f.Name(),
				filepath.Join(codeDir, f.Name()),
			})
		}
	}

	return services
}

func (is InstalledService) Cmd(c string) *exec.Cmd {
	cmd := exec.Command("make", c)
	cmd.Dir = is.Dir

	return cmd
}
