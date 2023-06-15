package security

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vimsucks/go-nexus-client/nexus3/pkg/client"
	"github.com/vimsucks/go-nexus-client/nexus3/pkg/tools"
	"github.com/vimsucks/go-nexus-client/nexus3/schema/security"
)

const (
	securityContentSelectorAPIEndpoint = securityAPIEndpoint + "/content-selectors"
)

type SecurityContentSelectorService client.Service

func NewSecurityContentSelectorService(c *client.Client) *SecurityContentSelectorService {

	s := &SecurityContentSelectorService{
		Client: c,
	}
	return s
}

func (s SecurityContentSelectorService) Create(cs security.ContentSelector) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(securityContentSelectorAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create content selector \"%s\": HTTP: %d, %s", cs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s SecurityContentSelectorService) List() ([]security.ContentSelector, error) {
	body, resp, err := s.Client.Get(securityContentSelectorAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read content selectors: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var contentSelectors []security.ContentSelector
	if err := json.Unmarshal(body, &contentSelectors); err != nil {
		return nil, fmt.Errorf("could not unmarshal content selector list: %v", err)
	}

	return contentSelectors, nil
}

func (s SecurityContentSelectorService) Get(name string) (*security.ContentSelector, error) {
	contentSelectors, err := s.List()
	if err != nil {
		return nil, err
	}

	for _, cs := range contentSelectors {
		if cs.Name == name {
			return &cs, nil
		}
	}

	return nil, nil
}

func (s SecurityContentSelectorService) Update(name string, cs security.ContentSelector) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", securityContentSelectorAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update content selector \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s SecurityContentSelectorService) Delete(name string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", securityContentSelectorAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete content selector \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
