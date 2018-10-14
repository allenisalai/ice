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
	"github.com/allenisalai/ice/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all installed services",
	Run: func(cmd *cobra.Command, args []string) {
		services := ice.GetInstalledServices(ice.GetConfig().CodeDir)

		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{"Name", "Path", "Running (coming soon)"}
		table.SetHeader(headers)

		for _, is := range services {
			row := []string{
				is.Name,
				is.Dir,
				"no",
			}
			table.Append(row)
		}
		table.Render() // Send output
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
