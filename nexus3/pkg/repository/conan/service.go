package conan

import (
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	conanAPIEndpoint = common.RepositoryAPIEndpoint + "/conan"
)

type RepositoryConanService struct {
	client *client.Client

	Proxy *RepositoryConanProxyService
}

func NewRepositoryConanService(c *client.Client) *RepositoryConanService {
	return &RepositoryConanService{
		client: c,

		Proxy: NewRepositoryConanProxyService(c),
	}
}
