package docker

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker"
	"golang.org/x/net/context"

	//"html"

	"strconv"
	"strings"
	"docker.io/go-docker/api/types/container"
	"io"
	"os"
)

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
