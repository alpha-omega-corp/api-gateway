package docker

import (
	"bytes"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bunrouter"
	"io"
	"mime/multipart"
	"net/http"
)

type CreateImageRequest struct {
	Dockerfile *multipart.FileHeader `form:"dockerfile"`
}

func CreateImageHandler(w http.ResponseWriter, req bunrouter.Request, s proto.DockerServiceClient) error {
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

	res, err := s.CreateImage(req.Context(), &proto.CreateImageRequest{
		Name:    req.FormValue("name"),
		Content: fileBuffer.Bytes(),
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
