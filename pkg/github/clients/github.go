package clients

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/routes"
	proto "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/config"
	"github.com/uptrace/bunrouter"
	"google.golang.org/grpc"
	"net/http"
)

type GitServiceClient interface {
	GetSecretContent(w http.ResponseWriter, req bunrouter.Request) error
	SyncEnvironment(w http.ResponseWriter, req bunrouter.Request) error
	DeleteSecret(w http.ResponseWriter, req bunrouter.Request) error
	CreateSecret(w http.ResponseWriter, req bunrouter.Request) error
	GetSecrets(w http.ResponseWriter, req bunrouter.Request) error
}

type gitClient struct {
	GitServiceClient

	client proto.GithubServiceClient
}

func NewGithubService(c *config.Host) GitServiceClient {
	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &gitClient{client: proto.NewGithubServiceClient(conn)}
}

func (svc *gitClient) GetSecretContent(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetSecretContentHandler(w, req, svc.client)
}

func (svc *gitClient) SyncEnvironment(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.SyncEnvironmentHandler(w, req, svc.client)
}

func (svc *gitClient) DeleteSecret(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeleteSecretHandler(w, req, svc.client)
}

func (svc *gitClient) CreateSecret(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreateSecretHandler(w, req, svc.client)
}

func (svc *gitClient) GetSecrets(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetSecretsHandler(w, req, svc.client)
}
