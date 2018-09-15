package cmd

import (
	"fmt"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/sysreq"
	"github.com/spf13/cobra"
	"log"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "View configuration and installed services",
	Long:  `View configuration and installed services`,
	Run: func(cmd *cobra.Command, args []string) {
		c := ice.GetConfig()
		log.Println(fmt.Sprintf("Log File: %s", c.LogFile))
		log.Println(fmt.Sprintf("Code Directory: %s", c.CodeDir))

		missingReqs := getMissingRequirements()

		if len(missingReqs) == 0 {
			log.Println("All system requirements are installed")
		} else {
			log.Println("Missing the following system requirements:")
			log.Println("-------------------------------------------")
			for i, mr := range missingReqs {
				log.Printf("%v) %s - %s", i+1, mr.Name(), mr.InstallationDescription())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func getMissingRequirements() []sysreq.SystemRequirementInterface {
	sysreq.AddRequirement(sysreq.CreateGitReq())
	sysreq.AddRequirement(sysreq.CreateDockerReq())

	return sysreq.CheckSystemRequirements()
}
