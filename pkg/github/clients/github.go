package clients

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/github/routes"
	proto "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/types"
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
	GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error
	DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	GetPackages(w http.ResponseWriter, req bunrouter.Request) error
	GetPackage(w http.ResponseWriter, req bunrouter.Request) error
	GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error
	CreatePackage(w http.ResponseWriter, req bunrouter.Request) error
	DeletePackage(w http.ResponseWriter, req bunrouter.Request) error
}

type gitClient struct {
	GitServiceClient

	client proto.GithubServiceClient
}

func NewGithubService(c types.ConfigHost) GitServiceClient {
	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &gitClient{client: proto.NewGithubServiceClient(conn)}
}

func (svc *gitClient) DeletePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageHandler(w, req, svc.client)
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

func (svc *gitClient) GetPackageTags(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageTagsHandler(w, req, svc.client)
}

func (svc *gitClient) DeletePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.DeletePackageVersionHandler(w, req, svc.client)
}

func (svc *gitClient) GetPackages(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackagesHandler(w, req, svc.client)
}

func (svc *gitClient) GetPackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageHandler(w, req, svc.client)
}

func (svc *gitClient) GetPackageFile(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.GetPackageFileHandler(w, req, svc.client)
}

func (svc *gitClient) CreatePackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageVersionHandler(w, req, svc.client)
}

func (svc *gitClient) PushPackageVersion(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.PushPackageVersionHandler(w, req, svc.client)
}

func (svc *gitClient) CreatePackage(w http.ResponseWriter, req bunrouter.Request) error {
	return routes.CreatePackageHandler(w, req, svc.client)
}
