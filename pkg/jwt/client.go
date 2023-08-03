package jwt

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/proto"
	"github.com/alpha-omega-corp/services/config"
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
