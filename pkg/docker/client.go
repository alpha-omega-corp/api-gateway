package docker

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/proto"
	"github.com/alpha-omega-corp/services/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.DockerServiceClient
}

func NewClient(c *config.Config) proto.DockerServiceClient {
	cc, err := grpc.Dial(c.DOCKER, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewDockerServiceClient(cc)
}
