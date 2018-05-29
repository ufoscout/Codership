package configuration

import (
	"fmt"
	"github.com/ufoscout/go-up"
)

/*
Load the configuration from the path folder
*/
func Load(configFile string) Config {

	up, err := go_up.NewGoUp().
		AddFile(configFile, false).
		AddReader(go_up.NewEnvReader("", false, false)). // Loading environment variables
		Build()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var config Config = Config{
		Server: ServerConfig{
			Port: up.GetInt("server.port"),
		},
		Frontend: FrontendConfig{
			ResourcesPath: up.GetString("frontend.resourcesPath"),
		},
	}

	return config
}
