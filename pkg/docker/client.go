package docker

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/types"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type Client interface {
	GetImage(w http.ResponseWriter, req bunrouter.Request) error
	StoreImage(w http.ResponseWriter, req bunrouter.Request) error
	BuildImage(w http.ResponseWriter, req bunrouter.Request) error
}

type dockerClient struct {
	Client
	client proto.DockerServiceClient
}

func NewClient(c types.ConfigHost) Client {
	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &dockerClient{client: proto.NewDockerServiceClient(conn)}
}

func (svc *dockerClient) GetImage(w http.ResponseWriter, req bunrouter.Request) error {
	return GetImageHandler(w, req, svc.client)
}

func (svc *dockerClient) StoreImage(w http.ResponseWriter, req bunrouter.Request) error {
	return StoreImageHandler(w, req, svc.client)
}

func (svc *dockerClient) BuildImage(w http.ResponseWriter, req bunrouter.Request) error {
	return BuildImageHandler(w, req, svc.client)
}
