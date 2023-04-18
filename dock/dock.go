package dock

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)
var (
    ctx    context.Context
	cli *client.Client
	err error
)

func InitDock(){
	ctx = context.Background()
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}
func PullImage(image string)error{

	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		fmt.Println("Could not pull image.")
		return err
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
	return nil
}
func CreateNewContainer(image, _HOST_PORT, CONTAINER_PORT string) (string, error) {
	err = PullImage(image)
	if err != nil {
		return "Image pull error.", err
	}


	config := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			nat.Port(CONTAINER_PORT+"/tcp"): struct{}{},
		},
	}
	
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(CONTAINER_PORT+"/tcp"): []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: _HOST_PORT,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil,nil, "")
	if err != nil {
		return "Container create error.", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "Container start error.", err
	}

	return resp.ID, nil
}

func GetAllContainers()([]types.Container, error){
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return []types.Container{}, err
	}

	return containers,nil
}
func StopOneContainer(containerID string)error{

	noWaitTimeout := 0 // to not wait for the container to exit gracefully
	if err := cli.ContainerStop(ctx, containerID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
		fmt.Println("Container ID:", containerID, "could not be stopped.")
		return err
	}
	return nil
}

func StopAllContainers()error{

	containers, err := GetAllContainers()
	if err != nil {
		fmt.Println("Containers could not be retreived.")
		cli.Close()
		return err 
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... \n")
		noWaitTimeout := 0 // to not wait for the container to exit gracefully
		if err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
			fmt.Println("Container ID:", container.ID, "could not be stopped.")
			return err
		}
	}
	cli.Close()
	return nil
}
func RemoveOneContainer(containerID string)error{
	if err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Println("Container ID:", containerID, "could not be stopped.")
		return err
	}
	return nil
}

func RestartContainer(containerID string)error{

	err := cli.ContainerRestart(ctx, containerID, container.StopOptions{})
	if err != nil{
		fmt.Println("Could not restart container", containerID)
		return err
	}

	return nil
}
func GetContainerStats(containerID string){	
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		fmt.Println("Could not get container logs. Err:", err)
	}

	io.Copy(os.Stdout, out)
}