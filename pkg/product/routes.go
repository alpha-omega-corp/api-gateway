package product

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/config"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt"
	"github.com/alpha-omega-corp/api-gateway/pkg/product/routes"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config, auth *jwt.ServiceClient) {
	m := jwt.NewAuthMiddleware(auth)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	cR := r.NewGroup("/product").Use(m.Auth)

	cR.GET("/", svc.Create)
}

func (svc *ServiceClient) Create(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateHandler(w, req, svc.Client)
}
