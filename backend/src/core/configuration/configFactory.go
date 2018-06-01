package configuration

import (
	"fmt"
	"github.com/ufoscout/go-up"
)

/*
LoadConfig the configuration from the path folder
*/
func LoadConfig(configFile string) Config {

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
			ResourcesPath: up.GetString("server.resourcesPath"),
		},
		Docker: DockerConfig{
			MariaDbImage: up.GetString("docker.images.mariadb"),
			MySqlImage: up.GetString("docker.images.mysql"),
		},
	}

	return config
}
