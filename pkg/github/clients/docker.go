package clients

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/routes"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type DockerServiceClient interface {
	GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error
	DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error
	GetPackages(w http.ResponseWriter, req bunrouter.Request) error
	GetPackage(w http.ResponseWriter, req bunrouter.Request) error
	GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackageContainer(w http.ResponseWriter, req bunrouter.Request) error
	GetContainerLogs(w http.ResponseWriter, req bunrouter.Request) error
	GetContainers(w http.ResponseWriter, req bunrouter.Request) error
	DeleteContainer(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackage(w http.ResponseWriter, req bunrouter.Request) error
}

type dockerService struct {
	DockerServiceClient

	client protoDocker.DockerServiceClient
}

func NewDockerService(c *config.Host) DockerServiceClient {
	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &dockerService{client: protoDocker.NewDockerServiceClient(conn)}
}

func (svc *dockerService) GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageTagsHandler(w, req, svc.client)
}

func (svc *dockerService) DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageVersionHandler(w, req, svc.client)
}

func (svc *dockerService) GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageVersionContainersHandler(w, req, svc.client)
}

func (svc *dockerService) GetPackages(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackagesHandler(w, req, svc.client)
}

func (svc *dockerService) GetPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageHandler(w, req, svc.client)
}

func (svc *dockerService) GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageFileHandler(w, req, svc.client)
}

func (svc *dockerService) CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageVersionHandler(w, req, svc.client)
}

func (svc *dockerService) PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.PushPackageVersionHandler(w, req, svc.client)
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

func (svc *dockerService) CreatePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageHandler(w, req, svc.client)
}
