package deployer

import (
	"github.com/gin-gonic/gin"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"github.com/ufoscout/Codership/backend/src/deployer/web"
)

type deployerModule struct {
	dockerConfig configuration.DockerConfig
	server *gin.Engine
	web *web.RestController
}

func DeployerModule(
	dockerConfig configuration.DockerConfig,
	server *gin.Engine) *deployerModule {
	return &deployerModule{
		dockerConfig: dockerConfig,
		server: server,
		web: web.NewRestController(server, nil),
	}
}

func (module *deployerModule) Start() {
	module.web.Start()
}
