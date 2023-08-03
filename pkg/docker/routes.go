package docker

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/docker/routes"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt"
	"github.com/alpha-omega-corp/api-gateway/types"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config, auth *jwt.ServiceClient) {
	m := types.NewAuthMiddleware(auth)

	svc := &ServiceClient{
		Client: Client(c),
	}

	cR := r.NewGroup("/docker").Use(m.Auth)
	cR.POST("/", svc.CreateContainer)
}

func (svc *ServiceClient) CreateContainer(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateContainerHandler(w, req, svc.Client)
}
