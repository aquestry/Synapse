package command

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/client"
)

type ListCommand struct{}

func (l *ListCommand) Execute(args []string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Docker client error:", err)
		return
	}
	cli.NegotiateAPIVersion(context.Background())

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Failed to list containers:", err)
		return
	}

	if len(containers) == 0 {
		fmt.Println("No containers found.")
		return
	}

	for _, c := range containers {
		fmt.Printf(" - %s (%s) [%s]\n", c.Names[0][1:], c.ID[:12], c.State)
	}
}
