package user

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/user/clients"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router, c *config.Host) clients.UserServiceClient {
	svc := clients.NewUserService(c)

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
