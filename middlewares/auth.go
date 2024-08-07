package middlewares

import (
	"errors"
	"github.com/alpha-omega-corp/api-gateway/pkg/user"
	"github.com/alpha-omega-corp/user-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	service user.Client
}

func NewAuthMiddleware(userService user.Client) *AuthMiddleware {
	return &AuthMiddleware{
		service: userService,
	}
}

func (middleware *AuthMiddleware) Auth(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		authHeader := req.Header.Get("Authorization")
		token := strings.Split(authHeader, "Bearer ")[1]

		res, err := middleware.service.Self().Validate(req.Context(), &proto.ValidateRequest{
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
