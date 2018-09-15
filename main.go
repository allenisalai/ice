package main

import (
	"github.com/allenisalai/ice/cmd"
	"github.com/allenisalai/ice/internal"
	"github.com/allenisalai/ice/internal/logger"
)

func main() {
	ice.InitializeAppConfigs()
	logger.SetUpLogger(ice.GetConfig().LogFile, true)

	cmd.Execute()
}
