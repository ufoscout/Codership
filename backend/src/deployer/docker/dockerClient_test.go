package docker

import (
	"testing"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"path"
	"github.com/ufoscout/Codership/backend/src/util"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestMariaDbDeployment(t *testing.T) {
	config := configuration.LoadConfig(path.Join(util.MainFolderPath(), configuration.CONFIG_FILE_NAME))
	docker := NewDockerDeployer(config.Docker)

	clusterSize := 2

	result,err := docker.DeployCluster("", "mariadb", clusterSize, 3306)
	assert.Nil(t, err)
	assert.Equal(t, clusterSize, len(result))
	fmt.Printf("Create container [%s]", result[0])
	fmt.Printf("Create container [%s]", result[1])
}
