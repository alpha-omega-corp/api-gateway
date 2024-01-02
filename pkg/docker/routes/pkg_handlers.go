package routes

import (
	"bytes"
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
)

type PushPackageRequestBody struct {
	Tag        string `json:"tag"`
	VersionSHA string `json:"sha"`
}

type DeletePackageRequestBody struct {
	Tag string `json:"tag"`
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

func DeletePackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(DeletePackageRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.DeletePackage(req.Context(), &proto.DeletePackageRequest{
		Name: req.Params().ByName("name"),
		Tag:  data.Tag,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func PushPackageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
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
