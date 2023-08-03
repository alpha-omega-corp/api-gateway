package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

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

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
