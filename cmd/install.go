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
	"github.com/spf13/cobra"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/repository"
	"os/exec"
	"fmt"
	"path/filepath"
	"os"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a new service",
	Long: `Install a repository from any configured repository source.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := *ice.GetConfig()
		rep := repository.Factory(c)

		repoFound := false
		for _, r := range rep.GetRepositories() {
			if r.Name == args[0] {
				repoFound = true
				fmt.Printf("1) Git install from: %s\n", r.SshUrl)
				cmd := exec.Command("git", "clone", r.SshUrl, filepath.Join(c.CodeDir, r.Name))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("2) %s repository cloned successfully\n", r.Name)
				// run make install
				cmd = exec.Command("rm", "-rf", filepath.Join(c.CodeDir, r.Name))
				cmd.Run()
				fmt.Printf("%s removed\n", filepath.Join(c.CodeDir, r.Name))
			}
		}

		if !repoFound {
			fmt.Printf("Unable to find repository: %s \n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
