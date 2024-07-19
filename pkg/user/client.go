package user

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/user/routes"
	"github.com/alpha-omega-corp/services/types"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type Service interface {
	Client() proto.UserServiceClient
	Login(w http.ResponseWriter, req bunrouter.Request) error
	Register(w http.ResponseWriter, req bunrouter.Request) error

	GetUsers(w http.ResponseWriter, req bunrouter.Request) error
	CreateUser(w http.ResponseWriter, req bunrouter.Request) error
	UpdateUser(w http.ResponseWriter, req bunrouter.Request) error
	DeleteUser(w http.ResponseWriter, req bunrouter.Request) error
	AssignUser(w http.ResponseWriter, req bunrouter.Request) error

	GetUserPermissions(w http.ResponseWriter, req bunrouter.Request) error
	GetServicePermissions(w http.ResponseWriter, req bunrouter.Request) error
	CreatePermission(w http.ResponseWriter, req bunrouter.Request) error
	GetServices(w http.ResponseWriter, req bunrouter.Request) error

	CreateRole(w http.ResponseWriter, req bunrouter.Request) error

	GetRoles(w http.ResponseWriter, req bunrouter.Request) error
}

type userService struct {
	Service

	client proto.UserServiceClient
}

func NewUserService(c types.ConfigHost) Service {
	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &userService{client: proto.NewUserServiceClient(conn)}
}

func (svc *userService) Client() proto.UserServiceClient {
	return svc.client
}
func (svc *userService) Login(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.LoginHandler(w, req, svc.client)
}
func (svc *userService) Register(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.RegisterHandler(w, req, svc.client)
}
func (svc *userService) GetUsers(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUsersHandler(w, req, svc.client)
}
func (svc *userService) CreateUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateUserHandler(w, req, svc.client)
}
func (svc *userService) UpdateUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.UpdateUserHandler(w, req, svc.client)
}
func (svc *userService) DeleteUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteUserHandler(w, req, svc.client)
}
func (svc *userService) AssignUser(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.AssignUserHandler(w, req, svc.client)
}
func (svc *userService) GetUserPermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetUserPermissionsHandler(w, req, svc.client)
}
func (svc *userService) GetServicePermissions(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetServicePermissionsHandler(w, req, svc.client)
}
func (svc *userService) CreatePermission(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePermissionHandler(w, req, svc.client)
}
func (svc *userService) GetServices(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetServices(w, req, svc.client)
}
func (svc *userService) GetRoles(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetRolesHandler(w, req, svc.client)
}
func (svc *userService) CreateRole(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateRoleHandler(w, req, svc.client)
}
