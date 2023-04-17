package dock

import (
	"context"
	"fmt"
	
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CreateNewContainer(image, _HOST_IP, _HOST_PORT, CONTAINER_PORT string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create Docker client.")
		return "", err
	}

	hostBinding := nat.PortBinding{
		HostIP:   _HOST_IP,
		HostPort: _HOST_PORT,
	}
	containerPort, err := nat.NewPort("tcp", CONTAINER_PORT)
	if err != nil {
		fmt.Printf("Unable to map host port %v to container port %v.",_HOST_PORT, CONTAINER_PORT)
		return "", err
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		}, nil, nil, "")
	if err != nil {
		fmt.Println("Container creation failed.")
		return "", err
	}
	
	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s has started", cont.ID)
	return cont.ID, nil
}


// func StopContainer(containerID string) error {
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	err = cli.ContainerStop(context.Background(), containerID, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return err
// }