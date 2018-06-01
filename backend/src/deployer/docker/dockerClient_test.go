package docker

import (
	"testing"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"path"
	"github.com/ufoscout/Codership/backend/src/util"
	"github.com/stretchr/testify/assert"
)

func TestMariaDbDeployment(t *testing.T) {
	config := configuration.LoadConfig(path.Join(util.MainFolderPath(), configuration.CONFIG_FILE_NAME))

	docker := NewDockerDeployer(config.Docker)
	result,_ := docker.DeployCluster("", "mariadb", 1)
	assert.True(t, result)
}
