package docker

import (
	"github.com/uptrace/bunrouter"
)

func RegisterClient(svc Client, r *bunrouter.Router) Client {

	r.GET("/docker/image/:name", svc.GetImage)
	r.POST("/docker/image", svc.StoreImage)
	r.POST("/docker/image/build", svc.BuildImage)

	return svc
}
