package clients

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/user/routes"
	"github.com/alpha-omega-corp/auth-svc/proto"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type UserServiceClient interface {
	Client() proto.AuthServiceClient

	Login(w http.ResponseWriter, req bunrouter.Request) error
	Register(w http.ResponseWriter, req bunrouter.Request) error
	CreateRole(w http.ResponseWriter, req bunrouter.Request) error
	GetUserPermissions(w http.ResponseWriter, req bunrouter.Request) error
	GetPermissions(w http.ResponseWriter, req bunrouter.Request) error
	CreatePermissions(w http.ResponseWriter, req bunrouter.Request) error
	GetPermissionServices(w http.ResponseWriter, req bunrouter.Request) error
	GetUsers(w http.ResponseWriter, req bunrouter.Request) error
	UpdateUser(w http.ResponseWriter, req bunrouter.Request) error
	AssignRole(w http.ResponseWriter, req bunrouter.Request) error
	GetRoles(w http.ResponseWriter, req bunrouter.Request) error
}

type userService struct {
	UserServiceClient

	client proto.AuthServiceClient
}

func NewUserService(c *config.Host) UserServiceClient {
	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &userService{client: proto.NewAuthServiceClient(conn)}
}

func (svc *userService) Client() proto.AuthServiceClient {
	return svc.client
}

func (svc *userService) GetUserPermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUserPermissionsHandler(w, req, svc.client)
}

func (svc *userService) GetPermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPermissionsHandler(w, req, svc.client)
}

func (svc *userService) CreatePermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePermissionsHandler(w, req, svc.client)
}

func (svc *userService) GetPermissionServices(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPermissionServicesHandler(w, req, svc.client)
}

func (svc *userService) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.client)
}

func (svc *userService) Register(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.Client())
}

func (svc *userService) CreateRole(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateRoleHandler(w, req, svc.Client())
}

func (svc *userService) GetRoles(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetRolesHandler(w, req, svc.Client())
}

func (svc *userService) AssignRole(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.AssignRoleHandler(w, req, svc.Client())

}

func (svc *userService) GetUsers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUsersHandler(w, req, svc.Client())
}

func (svc *userService) UpdateUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.UpdateUserHandler(w, req, svc.Client())
}
