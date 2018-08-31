package main

import (
	"github.com/allenisalai/ice/cmd"
	_ "github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal"
)

func main() {
	ice.InitializeAppConfigs()

	cmd.Execute()
}