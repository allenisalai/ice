package cmd

import (
	"fmt"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/repository"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// listCmd represents the list command
var includeDescription = false

const ARG_SEARCH_TERM = 0

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for repositories by name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := *ice.GetConfig()
		rep := repository.Factory(c)
		searchTerm := ""
		if len(args) > 0 {
			searchTerm = args[ARG_SEARCH_TERM]
		}

		repos, err := rep.GetRepositories()
		if err != nil {
			fmt.Println(err.Error())
		}

		var tableData [][]string
		for _, r := range repos {

			if searchTerm != "" && !strings.Contains(r.Name, searchTerm) {
				continue
			}

			row := []string{
				r.Name,
				r.Url,
			}
			if includeDescription {
				row = append(row, r.Description)
			}

			tableData = append(tableData, row)
		}

		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{"Name", "Url"}
		if includeDescription {
			headers = append(headers, "InstallationDescription")
		}
		table.SetHeader(headers)

		for _, v := range tableData {
			table.Append(v)
		}
		table.Render() // Send output

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&includeDescription, "description", "d", false, "Show the description of the of the repo")
}
