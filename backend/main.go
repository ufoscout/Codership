package main

import (
	"github.com/ufoscout/Codership/backend/src/core"
	"github.com/ufoscout/Codership/backend/src/configuration"
	)

func main() {

	config := configuration.LoadConfig(configuration.CONFIG_FILE_NAME)

	coreModule := core.CoreModule(&config)

	coreModule.Start()

}
