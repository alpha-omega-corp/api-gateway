package user

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/user/routes"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Host) *ServiceClient {
	svc := &ServiceClient{
		Client: NewClient(c),
	}

	r.POST("/login", svc.Login)
	r.POST("/register", svc.Register)

	r.GET("/roles", svc.GetRoles)
	r.POST("/role", svc.CreateRole)
	r.POST("/role/assign", svc.AssignRole)

	r.GET("/users", svc.GetUsers)
	r.POST("/user/:id", svc.UpdateUser)

	r.GET("/permission/pkg", svc.GetPermissionServices)

	r.POST("/permission", svc.CreatePermissions)
	r.GET("/permission/:serviceId", svc.GetPermissions)
	r.GET("/user/:id/permissions", svc.GetUserPermissions)

	return svc
}

func (svc *ServiceClient) GetUserPermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUserPermissionsHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPermissionsHandler(w, req, svc.Client)
}

func (svc *ServiceClient) CreatePermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePermissionsHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetPermissionServices(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPermissionServicesHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.Client)
}

func (svc *ServiceClient) Register(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.Client)
}

func (svc *ServiceClient) CreateRole(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateRoleHandler(w, req, svc.Client)
}

func (svc *ServiceClient) GetRoles(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetRolesHandler(w, req, svc.Client)
}

func (svc *ServiceClient) AssignRole(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.AssignRoleHandler(w, req, svc.Client)

}

func (svc *ServiceClient) GetUsers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUsersHandler(w, req, svc.Client)
}

func (svc *ServiceClient) UpdateUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.UpdateUserHandler(w, req, svc.Client)
}
