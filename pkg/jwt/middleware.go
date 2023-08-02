package jwt

import (
	"context"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt/pb"
	"github.com/alpha-omega-corp/api-gateway/pkg/utils"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	svc *ServiceClient
}

func NewAuthMiddleware(svc *ServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		svc: svc,
	}
}

func (m *AuthMiddleware) Auth(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		header := req.Header.Get("Authorization")

		if header == "" {
			return utils.ErrNotFound
		}

		token := strings.Split(header, "Bearer ")[1]

		res, err := m.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token,
		})

		if err != nil {
			return err
		}

		req.Context().Value(res.UserId)

		return next(w, req)
	}
}
