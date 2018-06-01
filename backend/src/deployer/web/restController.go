package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
)

type RestController struct {
	server *gin.Engine
	deployer common.Deployer
}

func NewRestController(
	server *gin.Engine,
	deployer common.Deployer) *RestController {
		return &RestController {
			server: server,
			deployer: deployer,
		}
}

func (web *RestController) Start() {

}
