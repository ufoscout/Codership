package docker

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker"
	"golang.org/x/net/context"

	"strconv"
	"strings"
	"docker.io/go-docker/api/types/container"
	"io"
	"os"
	"github.com/ufoscout/Codership/backend/src/configuration"
		"github.com/ufoscout/Codership/backend/src/deployer/common"
	"github.com/docker/go-connections/nat"
)

type dockerClient struct {
	dockerConfig configuration.DockerConfig
}

func NewDockerDeployer(dockerConfig configuration.DockerConfig) common.Deployer {
	return &dockerClient{
		dockerConfig: dockerConfig,
	}
}

func (d *dockerClient) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]string, error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()

	dockerImage := d.dockerConfig.MariaDbImage
	if ("mysql" == dbType) {
		dockerImage = d.dockerConfig.MySqlImage
	}

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

	ids := []string{}
	for i := 0; i < clusterSize; i++ {
		id, err := d.startNode(cli, ctx, dockerImage, clusterName, i, firstHostPort+i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (d *dockerClient) startNode(cli *docker.Client, ctx context.Context, dockerImage string, clusterName string, nodeNumber int, hostPort int) (string, error) {

	firstNodeName := clusterName + "-node-0"
	thisNodeName := clusterName + "-node-" + strconv.Itoa(nodeNumber)

	envVars := []string{
		"MYSQL_ROOT_PASSWORD=test",
		"MYSQL_DATABASE=test",
		"MYSQL_USER=test",
		"MYSQL_PASSWORD=test",
		"WSREP_NODE_NAME=" + thisNodeName,
		"WSREP_CLUSTER_NAME=galera_cluster",
		"WSREP_CLUSTER_ADDRESS=gcomm://" + firstNodeName,
	}

	portsMapping := nat.PortMap{
		"3306/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: strconv.Itoa(hostPort),	},},
	}

	hostConfig := &container.HostConfig{
		PortBindings: portsMapping,
	}

	if 0 == nodeNumber {
		envVars = append(envVars, "WSREP_NEW_CLUSTER=1")
	} else {
		fmt.Printf("Set wait hosts for node %d\n", nodeNumber)
		envVars = append(envVars, "WAIT_HOSTS=" + firstNodeName + ":3306")
		//hostConfig = &container.HostConfig{
		//	PortBindings: portsMapping,
		//	Links: []string{firstNodeName + ":" + firstNodeName},
		//}
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

func (d *dockerClient) createNetwork(cli *docker.Client, ctx context.Context, clusterName string) (types.NetworkCreateResponse, error) {
	return cli.NetworkCreate(ctx, d.networkName(clusterName), types.NetworkCreate{
		CheckDuplicate: true,
	})
}

func (d *dockerClient) networkName(clusterName string) string {
	return "network-" + clusterName
}

func (d *dockerClient) ClusterStatus(clusterName string) (map[string]string, error) {
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

func (d *dockerClient) RemoveCluster(clusterName string) (bool, error) {

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


func StartImage() {
		ctx := context.Background()
		cli, err := docker.NewEnvClient()
		if err != nil {
			panic(err)
		}

		pull, err := cli.ImagePull(ctx, "alpine:3.6", types.ImagePullOptions{})
		if err != nil {
			panic(err)
		} else {
			io.Copy(os.Stdout, pull)
		}

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: "alpine:3.6",
			Cmd:   []string{"echo", "hello world"},
		}, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
		}

		out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)
}

func Images() {

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//List all images available locally
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	htmlOutput := "<html>"
	for _, image := range images {
		htmlOutput += image.ID + " | " + strconv.Itoa(int(image.Size)) + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Println(htmlOutput)
}

func Containers() {

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	//Iterate through all containers and display each container's properties
	//fmt.Println("Image ID | Repo Tags | Size")
	htmlOutput := "<html>"
	for _, container := range containers {
		htmlOutput += strings.Join(container.Names, ",") + " | " + container.Image + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Println(htmlOutput)
}

func Networks() {

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	//List all networks
	htmlOutput := "<html>"
	//fmt.Println("Network Name | ID")
	for _, network := range networks {
		htmlOutput += network.Name + " | " + network.ID + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Println(htmlOutput)

}
