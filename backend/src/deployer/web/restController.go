package web

import (
	"github.com/gin-gonic/gin"
			"github.com/ufoscout/Codership/backend/src/deployer/service"
	"net/http"
	"fmt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RestController struct {
	server *gin.Engine
	service service.DeployerService
}

func NewRestController(
	server *gin.Engine,
	service service.DeployerService) *RestController {
		return &RestController {
			server: server,
			service: service,
		}
}

func (web *RestController) Start() {

	v1 := web.server.Group("/api/v1/cluster/:deploymentType")
	{
		v1.POST("/", web.createCluster)
		v1.GET("/:clusterName", web.getClusterStatus)
		v1.DELETE("/:clusterName", web.deleteCluster)
	}

}

func (web *RestController) createCluster(c *gin.Context) {

	fmt.Println("CREATE CLUSTER")

	deploymentType := c.Param("deploymentType")
	deployer, err := web.service.GetDeployer(deploymentType)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	fmt.Println("CREATE CLUSTER 2")
	var dto CreateClusterDTO
	c.BindJSON(&dto)
	response, err := deployer.DeployCluster(dto.ClusterName, dto.DbType, dto.ClusterSize, dto.FirstHostPort)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	fmt.Println("CLUSTER CREATED")
	c.JSON(http.StatusOK, response)
}

func (web *RestController) getClusterStatus(c *gin.Context) {

	deploymentType := c.Param("deploymentType")
	deployer, err := web.service.GetDeployer(deploymentType)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	clusterName := c.Param("clusterName")
	status, err := deployer.ClusterStatus(clusterName)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func (web *RestController) deleteCluster(c *gin.Context) {

	deploymentType := c.Param("deploymentType")
	deployer, err := web.service.GetDeployer(deploymentType)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	clusterName := c.Param("clusterName")
	deleted, err := deployer.RemoveCluster(clusterName)
	if err!=nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": deleted,
	})
}

func errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse {
		Error: err.Error(),
	})
}