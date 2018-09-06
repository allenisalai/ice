package cmd

import (
	"github.com/spf13/cobra"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/repository"
	"os"
	"github.com/olekukonko/tablewriter"
)

// listCmd represents the list command
var includeDescription = false

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := *ice.GetConfig()
		rep := repository.Factory(c)

		var data [][]string

		for _, r := range rep.GetRepositories()  {
			row :=[]string{
				r.Name,
				r.Url,
			}
			if includeDescription {
				row = append(row, r.Description)
			}

			data = append(data, row)
		}

		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{"Name", "Url"}
		if includeDescription {
			headers = append(headers, "Description")
		}
		table.SetHeader(headers)

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&includeDescription, "description", "d", false, "Show the description of the of the repo")
}
