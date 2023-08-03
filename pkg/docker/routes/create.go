package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type CreateContainerRequestBody struct {
	ImageName string `json:"imageName"`
}

func CreateContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(CreateContainerRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreateContainer(context.Background(), &proto.CreateContainerRequest{
		ImageName: data.ImageName,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
