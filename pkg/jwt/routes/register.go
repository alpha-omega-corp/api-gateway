package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/pb"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, req bunrouter.Request, s pb.AuthServiceClient) error {
	data := new(RegisterRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.Register(context.Background(), &pb.RegisterRequest{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
