package routes

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type CreateContainerRequestBody struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func CreateContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(CreateContainerRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreateContainer(req.Context(), &proto.CreateContainerRequest{
		Name:  data.Name,
		Image: data.Image,
	})

	if err != nil {
		return err
	}

	fmt.Print(res)

	return bunrouter.JSON(w, res)
}
