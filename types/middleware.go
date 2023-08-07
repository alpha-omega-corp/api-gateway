package types

import (
	"errors"
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/auth"
	"github.com/alpha-omega-corp/auth-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	svc *auth.ServiceClient
}

func NewAuthMiddleware(svc *auth.ServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		svc: svc,
	}
}

func (m *AuthMiddleware) Auth(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		header := req.Header.Get("Authorization")

		fmt.Print(header)

		token := strings.Split(header, "Bearer ")[1]

		res, err := m.svc.Client.Validate(req.Context(), &proto.ValidateRequest{
			Token: token,
		})

		if err != nil {
			return err
		}

		if int(res.GetStatus()) == http.StatusForbidden {
			w.WriteHeader(http.StatusForbidden)
			return errors.New("invalid token")
		}

		return next(w, req)
	}
}
