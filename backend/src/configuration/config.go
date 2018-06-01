package configuration

type ServerConfig struct {
	Port int
	ResourcesPath string
}

type DockerConfig struct {
	MariaDbImage string
	MySqlImage string
}

type Config struct {
	Server   ServerConfig
	Docker DockerConfig
}
