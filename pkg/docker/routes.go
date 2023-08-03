package docker

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/docker/routes"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt"
	"github.com/alpha-omega-corp/api-gateway/types"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config, s *jwt.ServiceClient) {
	middleware := types.NewAuthMiddleware(s)

	svc := &ServiceClient{
		Client: NewClient(c),
	}

	cR := r.NewGroup("/docker").
		Use(middleware.Auth)

	cR.POST("/", svc.CreateContainer)
}

func (svc *ServiceClient) CreateContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateContainerHandler(w, req, svc.Client)
}
