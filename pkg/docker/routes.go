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

	r.POST("/package", svc.CreatePackage)
	r.GET("/package/:name", svc.GetPackage)
	r.GET("/package/:name/tags", svc.GetPackageTags)
	r.DELETE("/package/:name/*version", svc.DeletePackageVersion)
	r.POST("/package/:name/container/:tag", svc.CreatePackageContainer)
	r.GET("/package/:name/containers/:tag", svc.GetPackageVersionContainers)
	r.POST("/package/:name/push", svc.PushPackageVersion)

	r.POST("/package/:name", svc.CreatePackageVersion)

	r.GET("/package/:name/:tag/:file", svc.GetPackageFile)

	r.GET("/containers", svc.GetContainers)
	r.DELETE("/container/:id", svc.DeleteContainer)
	r.GET("/container/:id/logs", svc.GetContainerLogs)
}

func (svc *ServiceClient) GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageTagsHandler(w, req, svc.Client)
}

func (svc *ServiceClient) DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageVersionHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageVersionContainersHandler(w, req, svc.Client)
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

func (svc *ServiceClient) PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.PushPackageVersionHandler(w, req, svc.Client)
}

func (svc *ServiceClient) CreatePackageContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageContainerHandler(w, req, svc.Client)
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
