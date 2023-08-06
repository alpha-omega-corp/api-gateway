package auth

import (
	"github.com/alpha-omega-corp/api-gateway/config"
	"github.com/alpha-omega-corp/api-gateway/pkg/auth/routes"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: NewClient(c),
	}

	r.POST("/login", svc.Login)
	r.POST("/register", svc.Register)
	r.POST("/validate", svc.Validate)

	return svc
}

func (svc *ServiceClient) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Register(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Validate(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.ValidateHandler(w, req, svc.Client)
}
