package routes

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

type CreateUserRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequestBody struct {
	Name string `json:"name"`
}

type AssignUserRequestBody struct {
	UserId int64   `json:"userId"`
	Roles  []int64 `json:"roles"`
}

func GetUsersHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	res, err := s.GetUsers(req.Context(), &proto.GetUsersRequest{})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func CreateUserHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(CreateUserRequestBody)

	fmt.Print(req.Body)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreateUser(req.Context(), &proto.CreateUserRequest{
		Name:  data.Name,
		Email: data.Email,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func UpdateUserHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	userId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	data := new(UpdateUserRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.UpdateUser(req.Context(), &proto.UpdateUserRequest{
		Id:   userId,
		Name: data.Name,
	})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func DeleteUserHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	userId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.DeleteUser(req.Context(), &proto.DeleteUserRequest{Id: userId})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func AssignUserHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(AssignUserRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.AssignUser(req.Context(), &proto.AssignUserRequest{
		UserId: data.UserId,
		Roles:  data.Roles,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
