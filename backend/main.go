package main

import (
	"github.com/ufoscout/Codership/backend/src/core"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"github.com/ufoscout/Codership/backend/src/deployer"
)

func main() {

	config := configuration.LoadConfig(configuration.CONFIG_FILE_NAME)

	coreModule := core.CoreModule(&config)
	deploymentModule := deployer.DeployerModule(config.Docker, coreModule.Server())

	deploymentModule.Start()
	coreModule.Start()

}
