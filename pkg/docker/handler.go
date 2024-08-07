package docker

import (
	"bytes"
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
)

type CreateImageRequest struct {
	Dockerfile *multipart.FileHeader `form:"dockerfile"`
}

type BuildImageRequest struct {
	Name string `form:"name"`
}

func GetImageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	res, err := s.GetImage(req.Context(), &proto.GetImageRequest{
		Name: req.Param("name"),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func StoreImageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
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

	res, err := s.StoreImage(req.Context(), &proto.StoreImageRequest{
		Name:    req.FormValue("name"),
		Content: fileBuffer.Bytes(),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func BuildImageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
	data := new(BuildImageRequest)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.BuildImage(req.Context(), &proto.BuildImageRequest{
		Name: data.Name,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
