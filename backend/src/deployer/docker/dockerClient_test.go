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
	// Setup configuration
	config := configuration.LoadConfig(path.Join(util.MainFolderPath(), configuration.CONFIG_FILE_NAME))
	docker := NewDockerDeployer(config.Docker)

	clusterName := "cluster1"
	clusterSize := 2

	// Create cluster
	containerIDs,err := docker.DeployCluster(clusterName, "mariadb", clusterSize, 3306)
	assert.Nil(t, err)
	assert.Equal(t, clusterSize, len(containerIDs))
	for i:=0; i<clusterSize; i++ {
		fmt.Printf("Created container [%s]\n", containerIDs[i])
	}

	// Check cluster status
	status, err := docker.ClusterStatus(clusterName)
	assert.Nil(t, err)
	for k, v := range status {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	assert.Equal(t, clusterSize, len(status))

	for i:=0; i<clusterSize; i++ {
		assert.Equal(t, "running", status[containerIDs[i]])
	}

	// Remove cluster
	remove, err := docker.RemoveCluster(clusterName)
	assert.Nil(t, err)
	assert.True(t, remove)

	// Check cluster status after removal
	_, err = docker.ClusterStatus(clusterName)
	assert.NotNil(t, err)

}
