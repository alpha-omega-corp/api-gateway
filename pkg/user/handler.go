package user

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/user/clients"
	"github.com/alpha-omega-corp/services/config"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router) clients.UserServiceClient {
	env, err := config.NewHandler().Environment("user")
	if err != nil {
		panic(err)
	}

	svc := clients.NewUserService(env.Host)

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
