package jwt

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/config"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AUTH, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
