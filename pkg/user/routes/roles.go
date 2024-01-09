package routes

import (
	"encoding/json"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type CreateRoleRequestBody struct {
	Name string `json:"name"`
}

type AssignRoleRequestBody struct {
	UserId int64 `json:"user"`
	RoleId int64 `json:"role"`
}

func CreateRoleHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(CreateRoleRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreateRole(req.Context(), &proto.CreateRoleRequest{
		Name: data.Name,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetRolesHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	res, err := s.GetRoles(req.Context(), &proto.GetRolesRequest{})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func AssignRoleHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(AssignRoleRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.AssignRole(req.Context(), &proto.AssignRoleRequest{
		UserId: data.UserId,
		RoleId: data.RoleId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
