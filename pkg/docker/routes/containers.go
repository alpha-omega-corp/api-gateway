package routes

import (
	"bytes"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type CreateContainerRequestBody struct {
	Dockerfile *multipart.FileHeader `form:"dockerfile"`
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

func GetPackagesHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetPackages(req.Context(), &proto.GetPackagesRequest{})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetPackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	pkgId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.GetPackage(req.Context(), &proto.GetPackageRequest{
		Id: pkgId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func DeletePackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	pkgId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.DeletePackage(req.Context(), &proto.DeletePackageRequest{
		Id: pkgId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func PushPackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	pkgId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.PushPackage(req.Context(), &proto.PushPackageRequest{
		Id: pkgId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func ContainerPackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	pkgId, err := strconv.ParseInt(req.Params().ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	res, err := s.ContainerPackage(req.Context(), &proto.ContainerPackageRequest{
		Id: pkgId,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
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
