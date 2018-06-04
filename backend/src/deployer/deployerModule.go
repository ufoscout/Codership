package deployer

import (
	"github.com/gin-gonic/gin"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"github.com/ufoscout/Codership/backend/src/deployer/web"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"github.com/ufoscout/Codership/backend/src/deployer/ansible"
	"github.com/ufoscout/Codership/backend/src/deployer/docker"
	"github.com/ufoscout/Codership/backend/src/deployer/kubernates"
	"github.com/ufoscout/Codership/backend/src/deployer/service"
)

type deployerModule struct {
	dockerConfig configuration.DockerConfig
	server *gin.Engine
	web *web.RestController
}

func DeployerModule(
	dockerConfig configuration.DockerConfig,
	server *gin.Engine) *deployerModule {

		deploymentService := service.NewDeployerService(
			map[string]common.Deployer{
				"ansible": ansible.NewAnsibleDeployer(),
				"docker": docker.NewDockerDeployer(dockerConfig),
				"kubernates": kubernates.NewKubernatesDeployer(),
			},
		)

	return &deployerModule{
		dockerConfig: dockerConfig,
		server: server,
		web: web.NewRestController(server, deploymentService),
	}
}

func (module *deployerModule) Start() {
	module.web.Start()
}
