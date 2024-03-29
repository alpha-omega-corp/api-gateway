package routes

import (
	"encoding/json"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, req bunrouter.Request, s proto.UserServiceClient) error {
	data := new(LoginRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.Login(req.Context(), &proto.LoginRequest{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
