package clients

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/routes"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	"github.com/alpha-omega-corp/services/types"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type DockerServiceClient interface {
	GetContainers(w http.ResponseWriter, req bunrouter.Request) error
	GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackageContainer(w http.ResponseWriter, req bunrouter.Request) error
	GetContainerLogs(w http.ResponseWriter, req bunrouter.Request) error
	DeleteContainer(w http.ResponseWriter, req bunrouter.Request) error
	StartContainer(w http.ResponseWriter, req bunrouter.Request) error
	StopContainer(w http.ResponseWriter, req bunrouter.Request) error
}

type dockerService struct {
	DockerServiceClient

	client protoDocker.DockerServiceClient
}

func NewDockerService(c types.ConfigHost) DockerServiceClient {
	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &dockerService{client: protoDocker.NewDockerServiceClient(conn)}
}

func (svc *dockerService) StopContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.StopContainerHandler(w, req, svc.client)
}

func (svc *dockerService) StartContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.StartContainerHandler(w, req, svc.client)
}

func (svc *dockerService) GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageVersionContainersHandler(w, req, svc.client)
}

func (svc *dockerService) CreatePackageContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageContainerHandler(w, req, svc.client)
}

func (svc *dockerService) GetContainerLogs(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainerLogsHandler(w, req, svc.client)
}

func (svc *dockerService) GetContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainersHandler(w, req, svc.client)
}

func (svc *dockerService) DeleteContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteContainerHandler(w, req, svc.client)
}
