package docker

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker"
	"golang.org/x/net/context"

	"strconv"
		"docker.io/go-docker/api/types/container"
	"io"
	"os"
	"github.com/ufoscout/Codership/backend/src/configuration"
		"github.com/ufoscout/Codership/backend/src/deployer/common"
	"github.com/docker/go-connections/nat"
	"log"
)

type dockerDeployer struct {
	dockerConfig configuration.DockerConfig
}

func NewDockerDeployer(dockerConfig configuration.DockerConfig) common.Deployer {
	return &dockerDeployer{
		dockerConfig: dockerConfig,
	}
}

func (d *dockerDeployer) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) (common.Nodes, error) {
	log.Printf("Start new cluster deployment. Cluster name: [%s], dbType: [%s], size: [%d], first port: [%d]\n",
		clusterName, dbType, clusterSize, firstHostPort)

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}

	dockerImage := d.dockerConfig.MariaDbImage
	if ("mysql" == dbType) {
		dockerImage = d.dockerConfig.MySqlImage
	}
	log.Printf("Pull dockerImage [%s]\n", dockerImage)

	pull, err := cli.ImagePull(ctx, dockerImage, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	} else {
		io.Copy(os.Stdout, pull)
	}

	_, err = d.createNetwork(cli, ctx, clusterName)
	if err != nil {
		return nil, err
	}

	nodes := []common.Node{}
	for i := 0; i < clusterSize; i++ {
		port := firstHostPort+i
		id, err := d.startNode(cli, ctx, dockerImage, clusterName, i, port)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, common.NewNode(id, "starting", port))
	}

	return nodes, nil
}

func (d *dockerDeployer) startNode(cli *docker.Client, ctx context.Context, dockerImage string, clusterName string, nodeNumber int, hostPort int) (string, error) {

	firstNodeName := clusterName + "-node-0"
	thisNodeName := clusterName + "-node-" + strconv.Itoa(nodeNumber)

	envVars := []string{
		"MYSQL_ROOT_PASSWORD=test",
		"MYSQL_DATABASE=test",
		"MYSQL_USER=test",
		"MYSQL_PASSWORD=test",
		"WSREP_NODE_NAME=" + thisNodeName,
		"WSREP_CLUSTER_NAME=galera_cluster",
	}

	portsMapping := nat.PortMap{
		"3306/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: strconv.Itoa(hostPort),	},},
	}

	hostConfig := &container.HostConfig{
		PortBindings: portsMapping,
	}

	if 0 == nodeNumber {
		envVars = append(envVars, "WSREP_NEW_CLUSTER=1")
		envVars = append(envVars, "WSREP_NEW_CLUSTER=1")
		envVars = append(envVars, "WSREP_CLUSTER_ADDRESS=gcomm://")
	} else {
		fmt.Printf("Set wait hosts for node %d\n", nodeNumber)
		envVars = append(envVars, "WAIT_HOSTS=" + firstNodeName + ":3306")
		envVars = append(envVars, "WAIT_HOSTS_TIMEOUT=60")
		envVars = append(envVars, "WSREP_CLUSTER_ADDRESS=gcomm://" + firstNodeName)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: dockerImage,
		//Cmd:   []string{"echo", "hello world"},
		ExposedPorts: nat.PortSet{
			nat.Port("3306/tcp"): {},
			//nat.Port("10001/tcp"): {},
		},
		Env: envVars,
	}, hostConfig, nil, thisNodeName)
	if err != nil {
		return "", err
	}

	err = cli.NetworkConnect(ctx, d.networkName(clusterName), resp.ID, nil)
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	/*
		statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				return false, err
			}
		case <-statusCh:
		}
	*/
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", err
	}

	io.Copy(os.Stdout, out)

	return resp.ID, nil
}

func (d *dockerDeployer) createNetwork(cli *docker.Client, ctx context.Context, clusterName string) (types.NetworkCreateResponse, error) {
	networkName := d.networkName(clusterName)
	log.Printf("Create network [%s]\n", networkName)
	return cli.NetworkCreate(ctx, networkName, types.NetworkCreate{
		CheckDuplicate: true,
	})
}

func (d *dockerDeployer) networkName(clusterName string) string {
	return "network-" + clusterName
}

func (d *dockerDeployer) ClusterStatus(clusterName string) (map[string]string, error) {
	log.Printf("Check status of cluster [%s]\n", clusterName)

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}

	resp, err := cli.NetworkInspect(ctx, d.networkName(clusterName), types.NetworkInspectOptions{})
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	for k,_ := range resp.Containers {
		resp, err := cli.ContainerInspect(ctx, k)
		if err==nil {
			result[k] = resp.State.Status
		}
	}
	return result, nil
}

func (d *dockerDeployer) RemoveCluster(clusterName string) (bool, error) {
	log.Printf("Remove cluster [%s]\n", clusterName)
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		return false, err
	}

	status, err := d.ClusterStatus(clusterName)
	if err != nil {
		return false, err
	}

	for id,_ := range status {

		err = cli.ContainerKill(ctx, id, "")
		if err != nil {
			return false, err
		}

		statusCh, errCh := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				return false, err
			}
		case <-statusCh:
		}

		err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
			RemoveVolumes: true,
		})
		if err != nil {
			return false, err
		}

	}

	err = cli.NetworkRemove(ctx, d.networkName(clusterName))
	if err != nil {
		return false, err
	}

	return true, nil
}
