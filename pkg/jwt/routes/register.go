package routes

import (
	"encoding/json"
	"github.com/alpha-omega-corp/authentication-svc/pkg/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	data := new(RegisterRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.Register(req.Context(), &proto.RegisterRequest{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
