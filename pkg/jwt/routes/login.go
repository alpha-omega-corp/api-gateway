package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/authentication-svc/pkg/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	data := new(LoginRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.Login(context.Background(), &proto.LoginRequest{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
