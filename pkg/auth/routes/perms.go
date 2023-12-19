package routes

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/auth-svc/proto"
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

func GetPermissionServicesHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	res, err := s.GetPermServices(req.Context(), &proto.GetPermServicesRequest{})

	if err != nil {
		return err
	}
	fmt.Print(res)

	return bunrouter.JSON(w, res)
}

func CreatePermissionsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	data := new(CreatePermissionsRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreatePermissions(req.Context(), &proto.CreatePermissionRequest{
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

func GetPermissionsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
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
