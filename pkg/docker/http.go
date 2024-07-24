package docker

import (
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router) Service {
	env, err := config.NewHandler().Environment("docker")
	if err != nil {
		panic(err)
	}

	svc := NewDockerService(env.Host)

	r.POST("/docker/image", svc.CreateImage)

	return svc
}
