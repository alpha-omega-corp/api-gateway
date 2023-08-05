package authentication

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/config"
	"github.com/alpha-omega-corp/authentication-svc/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func NewClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AUTH, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewAuthServiceClient(cc)
}
