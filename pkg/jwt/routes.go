package jwt

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/routes"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: NewClient(c),
	}

	cR := r.NewGroup("/auth")

	cR.POST("/login", svc.Login)
	cR.POST("/register", svc.Register)

	return svc
}

func (svc *ServiceClient) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Register(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Validate(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.Client)
}
