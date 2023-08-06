package routes

import (
	"errors"
	"github.com/alpha-omega-corp/auth-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strings"
)

func ValidateHandler(w http.ResponseWriter, req bunrouter.Request, s proto.AuthServiceClient) error {
	header := req.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(http.StatusForbidden)
		return errors.New("no authorization header")
	}

	token := strings.Split(header, "Bearer ")[1]

	res, err := s.Validate(req.Context(), &proto.ValidateRequest{
		Token: token,
	})

	w.WriteHeader(int(res.Status))

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res.User)
}
