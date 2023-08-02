package jwt

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/config"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/routes"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	cR := r.NewGroup("/auth")
	cR.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.Client)
}
