package routes

import (
	"bytes"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
)

type CreateContainerRequestBody struct {
	Dockerfile *multipart.FileHeader `form:"dockerfile"`
}

type CreatePackageRequestBody struct {
	Tag string `json:"tag"`
}

func CreatePackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	multipartFile, handler, err := req.FormFile("dockerfile")
	if err != nil {
		return err
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(multipartFile)

	file, err := handler.Open()
	if err != nil {
		return err
	}

	fileBuffer := bytes.NewBuffer(make([]byte, handler.Size))
	if _, err := io.Copy(fileBuffer, file); err != nil {
		return err
	}

	res, err := s.CreatePackage(req.Context(), &proto.CreatePackageRequest{
		Dockerfile: fileBuffer.Bytes(),
		Workdir:    req.FormValue("workdir"),
		Tag:        req.FormValue("tag"),
	})

	return bunrouter.JSON(w, res)
}

func GetContainerLogsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetContainerLogs(req.Context(), &proto.GetContainerLogsRequest{
		ContainerId: req.Params().ByName("id"),
	})

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

func GetPackagesHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetPackages(req.Context(), &proto.GetPackagesRequest{})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
