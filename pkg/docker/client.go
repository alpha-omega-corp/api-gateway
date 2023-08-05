package docker

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/config"
	"github.com/alpha-omega-corp/docker-svc/proto"
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
