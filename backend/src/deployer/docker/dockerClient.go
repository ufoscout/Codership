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
	"github.com/docker/go-connections/nat"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
)

type dockerClient struct {
	dockerConfig configuration.DockerConfig
}

func NewDockerDeployer(dockerConfig configuration.DockerConfig) common.Deployer {
	return &dockerClient{
		dockerConfig: dockerConfig,
	}
}

func (d *dockerClient) DeployCluster(clusterName string, dbType string, instances int) (bool, error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()

	dockerImage := d.dockerConfig.MariaDbImage
	if ("mysql" == dbType) {
		dockerImage = d.dockerConfig.MySqlImage
	}

	pull, err := cli.ImagePull(ctx, dockerImage, types.ImagePullOptions{})
	if err != nil {
		return false, err
	} else {
		io.Copy(os.Stdout, pull)
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3306/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "3306",	},},
		},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: dockerImage,
		//Cmd:   []string{"echo", "hello world"},
		ExposedPorts: nat.PortSet{
			nat.Port("3306/tcp"): {},
			//nat.Port("10001/tcp"): {},
		},
		Env: []string{
			"WSREP_NEW_CLUSTER=1",
			"MYSQL_ROOT_PASSWORD=test",
			"MYSQL_DATABASE=test",
			"MYSQL_USER=test",
			"MYSQL_PASSWORD=test",
			"WSREP_NODE_NAME=node1",
			"WSREP_CLUSTER_NAME=galera_cluster",
			"WSREP_CLUSTER_ADDRESS=gcomm://node1",
		},
	}, hostConfig, nil, "")
	if err != nil {
		return false, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return false, err
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
		return false, err
	}

	io.Copy(os.Stdout, out)

	return true, nil
}

func (d *dockerClient) RemoveCluster(clusterName string) (bool, error) {
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
