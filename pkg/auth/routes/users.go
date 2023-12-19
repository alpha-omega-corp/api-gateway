package routes

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/auth-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

type UpdateUserRequestBody struct {
	Name  string  `json:"name"`
	Roles []int64 `json:"roles"`
}

func GetUserPermissionsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	userId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.GetUserPermissions(req.Context(), &proto.GetUserPermissionsRequest{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetUsersHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	res, err := s.GetUsers(req.Context(), &proto.GetUsersRequest{})

	fmt.Print(res)
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func UpdateUserHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	userId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	data := new(UpdateUserRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	fmt.Print(data.Roles)
	res, err := s.UpdateUser(req.Context(), &proto.UpdateUserRequest{
		Id:    userId,
		Name:  data.Name,
		Roles: data.Roles,
	})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
