package configuration

// FrontendConfig contains the Frontend configuration
type FrontendConfig struct {
	ResourcesPath string
}

type ServerConfig struct {
	Port int
}

type Config struct {
	Server   ServerConfig
	Frontend FrontendConfig
}
