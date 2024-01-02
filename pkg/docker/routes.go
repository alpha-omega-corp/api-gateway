package docker

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/auth"
	"github.com/alpha-omega-corp/api-gateway/pkg/docker/routes"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Host, s *auth.ServiceClient) {
	svc := &ServiceClient{
		Client: NewClient(c),
	}

	r.GET("/packages", svc.GetPackages)
	r.POST("/packages/container", svc.ContainerPackage)

	r.POST("/package", svc.CreatePackage)
	r.GET("/package/:name", svc.GetPackage)
	r.DELETE("/package/:name", svc.DeletePackage)

	r.POST("/package/:name", svc.CreatePackageVersion)
	r.POST("/package/:name/push", svc.PushPackage)

	r.GET("/package/:name/:tag/:file", svc.GetPackageFile)

	r.GET("/containers", svc.GetContainers)
	r.DELETE("/container/:id", svc.DeleteContainer)
	r.GET("/container/:id/logs", svc.GetContainerLogs)
}

func (svc *ServiceClient) DeletePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPackages(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackagesHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageFileHandler(w, req, svc.Client)
}

func (svc *ServiceClient) CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageVersionHandler(w, req, svc.Client)
}

func (svc *ServiceClient) PushPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.PushPackageHandler(w, req, svc.Client)
}

func (svc *ServiceClient) ContainerPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.ContainerPackageHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetContainerLogs(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainerLogsHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainersHandler(w, req, svc.Client)
}

func (svc *ServiceClient) DeleteContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteContainerHandler(w, req, svc.Client)
}

func (svc *ServiceClient) CreatePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageHandler(w, req, svc.Client)
}
