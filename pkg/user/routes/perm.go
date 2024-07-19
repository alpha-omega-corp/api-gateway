package routes

import (
	"encoding/json"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

type CreatePermissionsRequestBody struct {
	RoleID    int64 `json:"roleId"`
	ServiceID int64 `json:"serviceId"`
	CanRead   bool  `json:"canRead"`
	CanWrite  bool  `json:"canWrite"`
	CanManage bool  `json:"canManage"`
}

func CreatePermissionHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(CreatePermissionsRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreatePermission(req.Context(), &proto.CreatePermissionRequest{
		RoleId:    data.RoleID,
		ServiceId: data.ServiceID,
		CanRead:   data.CanRead,
		CanWrite:  data.CanWrite,
		CanManage: data.CanManage,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetServices(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	res, err := s.GetServices(req.Context(), &proto.GetServicesRequest{})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetServicePermissionsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	serviceId, err := strconv.ParseInt(req.Params().ByName("serviceId"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.GetPermissions(req.Context(), &proto.GetPermissionsRequest{
		ServiceId: serviceId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetUserPermissionsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
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
