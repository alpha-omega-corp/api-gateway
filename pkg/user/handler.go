package user

import (
	"github.com/alpha-omega-corp/services/config"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bunrouter"
)

func RegisterRoutes(r *bunrouter.Router) Service {
	env, err := config.NewHandler().Environment("user")
	if err != nil {
		panic(err)
	}

	svc := NewUserService(env.Host)

	r.POST("/login", svc.Login)
	r.POST("/register", svc.Register)

	r.GET("/roles", svc.GetRoles)
	r.POST("/role", svc.CreateRole)

	r.GET("/users", svc.GetUsers)
	r.POST("/user", svc.CreateUser)
	r.PUT("/user/:id", svc.UpdateUser)
	r.DELETE("/user/:id", svc.DeleteUser)
	r.POST("/user/assign", svc.AssignUser)

	r.POST("/permission", svc.CreatePermission)
	r.GET("/permission/:serviceId", svc.GetServicePermissions)
	r.GET("/user/:id/permissions", svc.GetUserPermissions)
	r.GET("/permission/pkg", svc.GetServices)

	return svc
}
