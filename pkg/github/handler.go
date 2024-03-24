package github

import (
	"github.com/alpha-omega-corp/api-gateway/middlewares"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/clients"
	user "github.com/alpha-omega-corp/api-gateway/pkg/user/clients"
	"github.com/alpha-omega-corp/services/config"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router, u user.UserServiceClient) {
	jwt := middlewares.NewAuthMiddleware(u)
	env, err := config.NewHandler().Environment("github")
	if err != nil {
		panic(err)
	}

	r.Use(jwt.Auth).WithGroup("/github", func(g *bunrouter.Group) {
		svc := clients.NewGithubService(env.Host)

		g.GET("/secrets", svc.GetSecrets)
		g.POST("/secrets", svc.CreateSecret)
		g.DELETE("/secrets/:name", svc.DeleteSecret)
		g.GET("/secrets/:name", svc.GetSecretContent)
		g.POST("/secrets/sync", svc.SyncEnvironment)

		g.GET("/packages", svc.GetPackages)
		g.POST("/packages", svc.CreatePackage)
		g.GET("/packages/:name", svc.GetPackage)
		g.DELETE("/packages/:name", svc.DeletePackage)
		g.POST("/packages/:name", svc.CreatePackageVersion)
		g.GET("/packages/:name/tags", svc.GetPackageTags)
		g.DELETE("/packages/:name/*version", svc.DeletePackageVersion)
		g.POST("/packages/:name/push", svc.PushPackageVersion)
		g.GET("/packages/:name/:tag/:file", svc.GetPackageFile)
	})

	r.Use(jwt.Auth).WithGroup("/docker", func(g *bunrouter.Group) {
		svc := clients.NewDockerService(env.Host)

		g.WithGroup("/packages", func(pG *bunrouter.Group) {
			pG.POST("/:name/containers/:tag", svc.CreatePackageContainer)
			pG.GET("/:name/containers/:tag", svc.GetPackageVersionContainers)
		})

		g.WithGroup("/containers", func(cG *bunrouter.Group) {
			cG.GET("/", svc.GetContainers)
			cG.DELETE("/:id", svc.DeleteContainer)
			cG.GET("/:id/logs", svc.GetContainerLogs)
			cG.POST("/:id/start", svc.StartContainer)
			cG.POST("/:id/stop", svc.StopContainer)
		})
	})
}
