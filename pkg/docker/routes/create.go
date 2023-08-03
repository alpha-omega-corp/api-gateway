package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/api-gateway/pkg/docker/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

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

type CreateContainerRequestBody struct {
	ImageName string `json:"imageName"`
}
