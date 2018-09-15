// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/logger"
	"github.com/allenisalai/ice/internal/repository"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a new service",
	Long:  `Install a repository from any configured repository source.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := *ice.GetConfig()
		rep := repository.Factory(c)

		repoFound := false
		repos, err := rep.GetRepositories()

		if err != nil {
			log.Fatalf("Failed to retrieve repositories: %s", err.Error())
		}

		for _, r := range repos {
			if r.Name == args[0] {
				repoFound = true
				repoFolder := filepath.Join(c.CodeDir, r.Name)

				cloneGitRepo(r, repoFolder)
				runRepoInstallCommand(repoFolder)
			}
		}

		if !repoFound {
			log.Printf("Unable to find repository: '%s'\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func cloneGitRepo(r repository.Repository, repoFolder string) {
	fmt.Printf("1) Git install from: %s\n", r.SshUrl)
	if _, err := os.Stat(repoFolder); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", r.SshUrl, repoFolder)
		out, err := cmd.CombinedOutput()
		logger.Println(string(out))
		if err != nil {
			logger.Fatalf("Git clone failed: %s", err.Error())
		}
		log.Printf("'%s' repository cloned successfully\n", r.Name)
	} else {
		log.Printf("Repository direcory already exists: %s", repoFolder)
	}
}

func runRepoInstallCommand(repoFolder string) {
	log.Printf("Running repo install commands\n")
	cmd := exec.Command("make", "install")
	cmd.Dir = repoFolder
	out, err := cmd.CombinedOutput()
	logger.Println(string(out))
	if err != nil {
		log.Fatalf("Repo install failed: %s\n", err.Error())
	}
}
