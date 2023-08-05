package types

import (
	"github.com/alpha-omega-corp/api-gateway/pkg/authentication"
	"github.com/alpha-omega-corp/authentication-svc/proto"
	"github.com/alpha-omega-corp/services/httputils"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	svc *authentication.ServiceClient
}

func NewAuthMiddleware(svc *authentication.ServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		svc: svc,
	}
}

func (m *AuthMiddleware) Auth(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		header := req.Header.Get("Authorization")
		if header == "" {
			return httputils.ErrNotFound
		}

		token := strings.Split(header, "Bearer ")[1]

		_, err := m.svc.Client.Validate(req.Context(), &proto.ValidateRequest{
			Token: token,
		})

		if err != nil {
			return err
		}

		return next(w, req)
	}
}
