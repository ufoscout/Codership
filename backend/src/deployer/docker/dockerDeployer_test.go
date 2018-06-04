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
	"time"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
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
	var nodes []common.Node
	firstPort := 3306

	// Create cluster
	{
		var err error
		nodes, err = docker.DeployCluster(clusterName, dbType, clusterSize, firstPort)
		assert.Nil(t, err)
		assert.Equal(t, clusterSize, len(nodes))
		for i := 0; i < clusterSize; i++ {
			fmt.Printf("Created container [%s]\n", nodes[i].Id)
			assert.NotNil(t, nodes[i].Id)
			assert.NotNil(t, nodes[i].Status)
			assert.Equal(t, firstPort+i, nodes[i].Port)
		}
	}

	// Check cluster status
	{
		time.Sleep(time.Second)
		status, err := docker.ClusterStatus(clusterName)
		assert.Nil(t, err)
		for k, v := range status {
			fmt.Printf("key[%s] value[%s]\n", k, v)
		}
		assert.Equal(t, clusterSize, len(status))

		for i:=0; i<clusterSize; i++ {
			assert.Equal(t, "running", status[nodes[i].Id])
		}
	}


	// check that at least one node is synced with the cluster
	{
		ctx := context.Background()
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		for !logsContain(cli, ctx, nodes[0].Id, "WSREP: Shifting JOINED -> SYNCED") {
			time.Sleep(250 * time.Millisecond)
		}
	}

	// Remove the cluster
	{
		remove, err := docker.RemoveCluster(clusterName)
		assert.Nil(t, err)
		assert.True(t, remove)
	}

	// Check cluster status after removal
	{
		_, err := docker.ClusterStatus(clusterName)
		assert.NotNil(t, err)
	}

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
