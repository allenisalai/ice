package cmd

import (
	"github.com/spf13/cobra"
	"github.com/allenisalai/ice/internal"
	"fmt"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "View configuration and installed services",
	Long: `View configuration and installed services`,
	Run: func(cmd *cobra.Command, args []string) {
		c := ice.GetConfig()

		fmt.Println(fmt.Sprintf("Log File: %s", c.LogFile))
		fmt.Println(fmt.Sprintf("Code Directory: %s", c.CodeDir))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
