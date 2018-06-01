package main

import (
	"github.com/ufoscout/Codership/backend/src/core"
	"github.com/ufoscout/Codership/backend/src/configuration"
	)

type Module interface {
	Start()
}

func main() {

	config := configuration.LoadConfig(configuration.CONFIG_FILE_NAME)

	coreModule := core.CoreModule(&config)

	coreModule.Start()

}
