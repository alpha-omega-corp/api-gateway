package github

import (
	"github.com/alpha-omega-corp/api-gateway/middlewares"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/clients"
	userClients "github.com/alpha-omega-corp/api-gateway/pkg/user/clients"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Host, user userClients.UserServiceClient) {
	jwt := middlewares.NewAuthMiddleware(user)

	r.Use(jwt.Auth).WithGroup("/github", func(g *bunrouter.Group) {
		svc := clients.NewGithubService(c)

		g.GET("/secrets", svc.GetSecrets)
		g.POST("/secrets", svc.CreateSecret)
		g.DELETE("/secrets/:name", svc.DeleteSecret)
		g.GET("/secrets/:name", svc.GetSecretContent)
		g.POST("/secrets/sync", svc.SyncEnvironment)
	})

	r.Use(jwt.Auth).WithGroup("/docker", func(g *bunrouter.Group) {
		svc := clients.NewDockerService(c)

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
