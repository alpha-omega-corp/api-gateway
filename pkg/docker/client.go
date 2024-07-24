package docker

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/types"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type Service interface {
	CreateImage(w http.ResponseWriter, req bunrouter.Request) error
}

type dockerService struct {
	Service
	client proto.DockerServiceClient
}

func NewDockerService(c types.ConfigHost) Service {
	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &dockerService{client: proto.NewDockerServiceClient(conn)}
}

func (svc *dockerService) CreateImage(w http.ResponseWriter, req bunrouter.Request) error {
	return CreateImageHandler(w, req, svc.client)
}
