package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project based off one of the configured skeleton projects.",
	Long: `Create a new project based off one of the configured skeleton projects.  
If you are not seeing any options, check your configuration file to ensure they are configured properly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TO BE COMPLETED")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
