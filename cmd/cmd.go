package cmd

import (
	"errors"
	"github.com/allenisalai/ice/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// cmdCmd represents the cmd command
var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Run a command on a given service `cmd [service] [command]`",
	Long: `Run a command on a given service,  Available commands are:
		start: Start a docker daemon of the service
		stop: Stop the docker daemon of the service
		wipe: Completely remove the docker volumes 
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatalln("You are missing required arguments [service] [command]")
		}
		serviceName := args[0]
		installedServices := ice.GetInstalledServices(ice.GetConfig().CodeDir)
		is, err := findInstalledService(serviceName, installedServices)
		if err != nil {
			log.Fatalf(err.Error())
		}

		serviceCommand := args[1]
		if isValidDevCommand(serviceCommand) == false {
			log.Fatalf("%s isn't a valid command.  start, stop, wipe athe the only valid commands", serviceCommand)
		}

		c := is.Cmd("dev" + serviceCommand)
		c.Stdout = os.Stdout
		err = c.Run()
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdCmd)
}

func isValidDevCommand(cmd string) bool {
	for _, c := range []string{"start", "stop", "wipe"} {
		if c == cmd {
			return true
		}
	}

	return false
}

func findInstalledService(sn string, is []ice.InstalledService) (ice.InstalledService, error) {
	for _, s := range is {
		if s.Name == sn {
			return s, nil
		}
	}

	return ice.InstalledService{}, errors.New(sn + " is not installed.")
}
