package routes

import (
	"bytes"
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type PushPackageRequestBody struct {
	Tag        string `json:"tag"`
	VersionSHA string `json:"sha"`
}

type DeletePackageRequestBody struct {
	Tag string `json:"tag"`
}

type CreatePackageContainerRequestBody struct {
	ContainerName string `json:"containerName"`
}

type GetPackageVersionContainers struct {
	Path string `json:"path"`
}

func GetPackageTagsHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetPackageTags(req.Context(), &proto.GetPackageTagsRequest{
		Name: req.Params().ByName("name"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func DeletePackageVersionHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	versions := strings.Split(req.Params().ByName("version"), "/")
	versionTag := versions[0]
	versionId, _ := strconv.ParseInt(versions[1], 10, 64)

	res, err := s.DeletePackageVersion(req.Context(), &proto.DeletePackageVersionRequest{
		Name:    req.Params().ByName("name"),
		Tag:     versionTag,
		Version: &versionId,
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

func GetPackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	name := req.Params().ByName("name")

	res, err := s.GetPackage(req.Context(), &proto.GetPackageRequest{
		Name: name,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetPackageFileHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	p := req.Params().ByName("name") + "/" + req.Params().ByName("tag")

	res, err := s.GetPackageFile(req.Context(), &proto.GetPackageFileRequest{
		Path: p,
		Name: req.Params().ByName("file"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func CreatePackageVersionHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	contents, handler, err := req.FormFile("content")

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(contents)

	file, err := handler.Open()
	if err != nil {
		return err
	}

	fileBuffer := bytes.NewBuffer(make([]byte, handler.Size))
	if _, err := io.Copy(fileBuffer, file); err != nil {
		return err
	}

	res, err := s.CreatePackageVersion(req.Context(), &proto.CreatePackageVersionRequest{
		Name:    req.Params().ByName("name"),
		Tag:     req.FormValue("tag"),
		Content: fileBuffer.Bytes(),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func PushPackageVersionHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(PushPackageRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.PushPackage(req.Context(), &proto.PushPackageRequest{
		Name:       req.Params().ByName("name"),
		Tag:        data.Tag,
		VersionSHA: data.VersionSHA,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
