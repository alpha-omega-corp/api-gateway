package routes

import (
	"encoding/json"
	"fmt"
	protoGithub "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type CreateSecretRequestBody struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func GetSecretContentHandler(w http.ResponseWriter, req bunrouter.Request, client protoGithub.GithubServiceClient) error {
	res, err := client.GetSecretContent(req.Context(), &protoGithub.GetSecretContentRequest{
		Name: req.Params().ByName("name"),
	})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func SyncEnvironmentHandler(w http.ResponseWriter, req bunrouter.Request, client protoGithub.GithubServiceClient) error {
	res, err := client.SyncEnvironment(req.Context(), &protoGithub.SyncEnvironmentRequest{})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func GetSecretsHandler(w http.ResponseWriter, req bunrouter.Request, client protoGithub.GithubServiceClient) error {
	res, err := client.GetSecrets(req.Context(), &protoGithub.GetSecretsRequest{})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func CreateSecretHandler(w http.ResponseWriter, req bunrouter.Request, client protoGithub.GithubServiceClient) error {
	data := new(CreateSecretRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	fmt.Print(data)
	res, err := client.CreateSecret(req.Context(), &protoGithub.CreateSecretRequest{
		Name:    data.Name,
		Content: []byte(data.Content),
	})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

func DeleteSecretHandler(w http.ResponseWriter, req bunrouter.Request, client protoGithub.GithubServiceClient) error {
	res, err := client.DeleteSecret(req.Context(), &protoGithub.DeleteSecretRequest{
		Name: req.Params().ByName("name"),
	})
	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}
