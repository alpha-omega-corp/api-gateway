package github

import (
	"github.com/alpha-omega-corp/api-gateway/middlewares"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/routes"
	"github.com/alpha-omega-corp/api-gateway/pkg/user"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Host, authSvc *user.ServiceClient) {
	authentication := middlewares.NewAuthMiddleware(authSvc)

	r.Use(authentication.Auth).WithGroup("/github", func(g *bunrouter.Group) {
		svc := &GitServiceClient{Client: NewGithubClient(c)}

		g.GET("/secrets", svc.GetSecrets)
		g.POST("/secrets", svc.CreateSecret)
		g.DELETE("/secrets/:name", svc.DeleteSecret)
		g.GET("/secrets/:name", svc.GetSecretContent)
		g.POST("/secrets/sync", svc.SyncEnvironment)
	})

	r.Use(authentication.Auth).WithGroup("/docker", func(g *bunrouter.Group) {
		svc := &DockerServiceClient{Client: NewDockerClient(c)}

		g.GET("/packages", svc.GetPackages)
		g.POST("/packages", svc.CreatePackage)
		g.GET("/packages/:name", svc.GetPackage)
		g.POST("/packages/:name", svc.CreatePackageVersion)
		g.GET("/packages/:name/tags", svc.GetPackageTags)
		g.DELETE("/packages/:name/*version", svc.DeletePackageVersion)
		g.POST("/packages/:name/container/:tag", svc.CreatePackageContainer)
		g.GET("/packages/:name/containers/:tag", svc.GetPackageVersionContainers)
		g.POST("/packages/:name/push", svc.PushPackageVersion)
		g.GET("/packages/:name/:tag/:file", svc.GetPackageFile)

		g.WithGroup("/containers", func(cG *bunrouter.Group) {
			cG.GET("/", svc.GetContainers)
			cG.DELETE("/:id", svc.DeleteContainer)
			cG.GET("/:id/logs", svc.GetContainerLogs)
		})
	})
}

func (svc *GitServiceClient) GetSecretContent(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetSecretContentHandler(w, req, svc.Client)
}

func (svc *GitServiceClient) SyncEnvironment(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.SyncEnvironmentHandler(w, req, svc.Client)
}

func (svc *GitServiceClient) DeleteSecret(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteSecretHandler(w, req, svc.Client)
}

func (svc *GitServiceClient) CreateSecret(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateSecretHandler(w, req, svc.Client)
}

func (svc *GitServiceClient) GetSecrets(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetSecretsHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageTagsHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageVersionHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetPackageVersionContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageVersionContainersHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetPackages(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackagesHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageFileHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageVersionHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.PushPackageVersionHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) CreatePackageContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageContainerHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetContainerLogs(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainerLogsHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) GetContainers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetContainersHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) DeleteContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteContainerHandler(w, req, svc.Client)
}

func (svc *DockerServiceClient) CreatePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageHandler(w, req, svc.Client)
}
