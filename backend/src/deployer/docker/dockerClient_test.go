package docker

import (
	"testing"
	"github.com/ufoscout/Codership/backend/src/configuration"
	"path"
	"github.com/ufoscout/Codership/backend/src/util"
	"github.com/stretchr/testify/assert"
	"fmt"
	"strings"
	client "docker.io/go-docker"
	"context"
	"docker.io/go-docker/api/types"
			"io/ioutil"
		)

func TestMariaDbDeployment(t *testing.T) {
	startDbDeployment(t, "mariadb")
}

func TestMysqlDeployment(t *testing.T) {
	startDbDeployment(t, "mysql")
}

func startDbDeployment(t *testing.T, dbType string) {
	// Setup configuration
	config := configuration.LoadConfig(path.Join(util.MainFolderPath(), configuration.CONFIG_FILE_NAME))
	docker := NewDockerDeployer(config.Docker)

	clusterName := "cluster-1"
	clusterSize := 2

	// Create cluster
	containerIDs,err := docker.DeployCluster(clusterName, dbType, clusterSize, 3306)
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

	// wait for at least one node to sync with the cluster
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	for !logsContain(cli, ctx, containerIDs[0], "WSREP: Shifting JOINED -> SYNCED") {}

	// Remove cluster
	remove, err := docker.RemoveCluster(clusterName)
	assert.Nil(t, err)
	assert.True(t, remove)

	// Check cluster status after removal
	_, err = docker.ClusterStatus(clusterName)
	assert.NotNil(t, err)

}

func logsContain(cli *client.Client, ctx context.Context, containerID string, substring string) bool {
	rc, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail: "all",
	})
	if err != nil {
		panic(err)
	}

	logs := ""
	if b, err := ioutil.ReadAll(rc); err == nil {
		logs = string(b)
	}

	return strings.Contains(logs, substring)
}
