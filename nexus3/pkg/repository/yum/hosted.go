package yum

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
	yumHostedAPIEndpoint = yumAPIEndpoint + "/hosted"
)

type RepositoryYumHostedService struct {
	client *client.Client
}

func NewRepositoryYumHostedService(c *client.Client) *RepositoryYumHostedService {
	return &RepositoryYumHostedService{
		client: c,
	}
}

func (s *RepositoryYumHostedService) Create(repo repository.YumHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(yumHostedAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryYumHostedService) Get(id string) (*repository.YumHostedRepository, error) {
	var repo repository.YumHostedRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", yumHostedAPIEndpoint, id), nil)
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

func (s *RepositoryYumHostedService) Update(id string, repo repository.YumHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", yumHostedAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryYumHostedService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
