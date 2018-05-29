package main

import (
	"github.com/ufoscout/Codership/backend/src/core"
	"github.com/ufoscout/Codership/backend/src/core/configuration"
	)

func main() {

	config := configuration.Load(configuration.CONFIG_FILE_NAME)

	coreModule := core.CoreModule(&config)

	coreModule.Start()

}
