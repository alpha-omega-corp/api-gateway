package docker

import (
	"github.com/alpha-omega-corp/api-gateway/config"
	"github.com/alpha-omega-corp/api-gateway/pkg/auth"
	"github.com/alpha-omega-corp/api-gateway/pkg/docker/routes"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config, s *auth.ServiceClient) {
	svc := &ServiceClient{
		Client: NewClient(c),
	}

	r.GET("/packages", svc.GetPackages)
	r.POST("/package", svc.CreatePackage)
	r.GET("/package/:id", svc.GetPackage)
	r.DELETE("/package/:id", svc.DeletePackage)
	r.POST("/package/:id/push", svc.PushPackage)
	r.POST("/package/:id/container", svc.ContainerPackage)

	r.GET("/containers", svc.GetContainers)
	r.GET("/container/:id/logs", svc.GetContainerLogs)
}

func (svc *ServiceClient) GetPackages(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackagesHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageHandler(w, req, svc.Client)
}

func (svc *ServiceClient) DeletePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageHandler(w, req, svc.Client)
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

func (svc *ServiceClient) CreatePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageHandler(w, req, svc.Client)
}
