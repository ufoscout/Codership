package web

import (
	"github.com/gin-gonic/gin"
			"github.com/ufoscout/Codership/backend/src/deployer/service"
)

type RestController struct {
	server *gin.Engine
	service *service.DeployerService
}

func NewRestController(
	server *gin.Engine,
	service *service.DeployerService) *RestController {
		return &RestController {
			server: server,
			service: service,
		}
}

func (web *RestController) Start() {

}
