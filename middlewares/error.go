package middlewares

import (
	"github.com/alpha-omega-corp/services/httputils"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func NewErrorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err == nil {
			return nil
		}

		httpErr := httputils.From(err, true)
		if httpErr.Status != 0 {
			w.WriteHeader(httpErr.Status)
		}
		_ = bunrouter.JSON(w, httpErr)

		return err
	}
}
