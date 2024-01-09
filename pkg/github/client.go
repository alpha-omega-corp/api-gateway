package github

import (
	"fmt"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	protoGithub "github.com/alpha-omega-corp/github-svc/proto/github"

	"github.com/alpha-omega-corp/services/config"
	"google.golang.org/grpc"
)

type DockerServiceClient struct {
	Client protoDocker.DockerServiceClient
}

type GitServiceClient struct {
	Client protoGithub.GithubServiceClient
}

func NewDockerClient(c *config.Host) protoDocker.DockerServiceClient {
	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return protoDocker.NewDockerServiceClient(conn)
}

func NewGithubClient(c *config.Host) protoGithub.GithubServiceClient {
	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return protoGithub.NewGithubServiceClient(conn)
}
