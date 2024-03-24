package routes

import (
	"encoding/json"
	"fmt"
	proto "github.com/alpha-omega-corp/github-svc/proto/docker"
	"github.com/uptrace/bunrouter"
	"mime/multipart"
	"net/http"
)

type CreateContainerRequestBody struct {
	Dockerfile *multipart.FileHeader `form:"dockerfile"`
}

func StopContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.StopContainer(req.Context(), &proto.StopContainerRequest{
		ContainerId: req.Params().ByName("id"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func StartContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.StartContainer(req.Context(), &proto.StartContainerRequest{
		ContainerId: req.Params().ByName("id"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetContainerLogsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetContainerLogs(req.Context(), &proto.GetContainerLogsRequest{
		ContainerId: req.Params().ByName("id"),
	})

	fmt.Print(res)
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetContainersHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetContainers(req.Context(), &proto.GetContainersRequest{})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func DeleteContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.DeleteContainer(req.Context(), &proto.DeleteContainerRequest{
		ContainerId: req.Params().ByName("id"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func CreatePackageContainerHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(CreatePackageContainerRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	path := req.Params().ByName("name") + "/" + req.Params().ByName("tag")
	res, err := s.CreatePackageContainer(req.Context(), &proto.CreatePackageContainerRequest{
		Path: path,
		Name: data.ContainerName,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetPackageVersionContainersHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetPackageVersionContainers(req.Context(), &proto.GetPackageVersionContainersRequest{
		Path: req.Params().ByName("name") + "/" + req.Params().ByName("tag"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
