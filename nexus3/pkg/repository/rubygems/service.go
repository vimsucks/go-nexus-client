package rubygems

import (
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	rubyGemsAPIEndpoint = common.RepositoryAPIEndpoint + "/rubygems"
)

type RepositoryRubyGemsService struct {
	client *client.Client

	Group  *RepositoryRubyGemsGroupService
	Hosted *RepositoryRubyGemsHostedService
	Proxy  *RepositoryRubyGemsProxyService
}

func NewRepositoryRubyGemsService(c *client.Client) *RepositoryRubyGemsService {
	return &RepositoryRubyGemsService{
		client: c,

		Group:  NewRepositoryRubyGemsGroupService(c),
		Hosted: NewRepositoryRubyGemsHostedService(c),
		Proxy:  NewRepositoryRubyGemsProxyService(c),
	}
}
