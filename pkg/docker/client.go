package docker

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.DockerServiceClient
}

func NewClient(c *config.Host) proto.DockerServiceClient {
	cc, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewDockerServiceClient(cc)
}
