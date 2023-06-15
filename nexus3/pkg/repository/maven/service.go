package maven

import (
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	mavenAPIEndpoint = common.RepositoryAPIEndpoint + "/maven"
)

type RepositoryMavenService struct {
	client *client.Client

	Group  *RepositoryMavenGroupService
	Hosted *RepositoryMavenHostedService
	Proxy  *RepositoryMavenProxyService
}

func NewRepositoryMavenService(c *client.Client) *RepositoryMavenService {
	return &RepositoryMavenService{
		client: c,

		Group:  NewRepositoryMavenGroupService(c),
		Hosted: NewRepositoryMavenHostedService(c),
		Proxy:  NewRepositoryMavenProxyService(c),
	}
}
