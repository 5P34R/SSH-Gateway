package gateway

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerAddress() []string {

	containerAddList := []string{}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	cli.NegotiateAPIVersion(ctx)

	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: false,
	})
	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	for _, container := range containers {

		port := container.Ports[0].PublicPort
		fmt.Println("localhost:" + strconv.FormatUint(uint64(port), 10))
		containerAddList = append(containerAddList, "localhost:"+strconv.FormatUint(uint64(port), 10))
	}

	return containerAddList
}

func GetNumberOfContainer() int {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	cli.NegotiateAPIVersion(ctx)
	if err != nil {
		log.Fatalf("Failed to create docker client: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: false,
	})

	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	return len(containers)
}

func ContainerList() []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	cli.NegotiateAPIVersion(ctx)
	if err != nil {
		log.Fatalf("Failed to create docker client: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: false,
	})

	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	return containers
}
