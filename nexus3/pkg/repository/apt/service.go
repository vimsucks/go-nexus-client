package apt

import (
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	repositoryAptAPIEndpoint = common.RepositoryAPIEndpoint + "/apt"
)

type RepositoryAptService struct {
	client *client.Client

	Hosted *RepositoryAptHostedService
	Proxy  *RepositoryAptProxyService
}

func NewRepositoryAptService(c *client.Client) *RepositoryAptService {
	return &RepositoryAptService{
		client: c,

		Hosted: NewRepositoryAptHostedService(c),
		Proxy:  NewRepositoryAptProxyService(c),
	}
}
