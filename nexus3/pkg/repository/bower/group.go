package bower

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/tools"
	"github.com/vimsucks/go-nexus-client/nexus3/schema/repository"
)

const (
	bowerGroupAPIEndpoint = repositoryBowerAPIEndpoint + "/group"
)

type RepositoryBowerGroupService struct {
	client *client.Client
}

func NewRepositoryBowerGroupService(c *client.Client) *RepositoryBowerGroupService {
	return &RepositoryBowerGroupService{
		client: c,
	}
}

func (s *RepositoryBowerGroupService) Create(repo repository.BowerGroupRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(bowerGroupAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryBowerGroupService) Get(id string) (*repository.BowerGroupRepository, error) {
	var repo repository.BowerGroupRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", bowerGroupAPIEndpoint, id), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &repo); err != nil {
		return nil, fmt.Errorf("could not unmarshal repository: %v", err)
	}
	return &repo, nil
}

func (s *RepositoryBowerGroupService) Update(id string, repo repository.BowerGroupRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", bowerGroupAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryBowerGroupService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
