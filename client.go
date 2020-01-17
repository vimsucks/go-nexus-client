package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Client represents the Nexus API Client interface
type Client interface {
	ContentTypeTextPlain()
	ContentTypeJSON()
	RepositoryCreate(Repository, string, string) error
	RepositoryRead(string, string, string) (*Repository, error)
	RepositoryUpdate(string, Repository) error
	RepositoryDelete(string) error
	RoleCreate(Role) error
	RoleRead(string) (*Role, error)
	RoleUpdate(string, Role) error
	RoleDelete(string) error
	UserCreate(User) error
	UserRead(string) (*User, error)
	UserUpdate(string, User) error
	UserDelete(string) error
	UserChangePassword(string, string) error
}

type client struct {
	config      Config
	contentType string
	client      *http.Client
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config Config) Client {
	return &client{
		config:      config,
		contentType: "application/json",
		client:      &http.Client{},
	}
}

func (c client) setContentType(s string) {
	log.Printf("Setting ContentType: %s", s)
	c.contentType = s
}

func (c client) ContentTypeJSON() {
	c.setContentType("application/json")
}

func (c client) ContentTypeTextPlain() {
	c.setContentType("text/plain")
}

func (c client) NewRequest(method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	url := fmt.Sprintf("%s/%s", c.config.URL, endpoint)
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	req.SetBasicAuth(c.config.Username, c.config.Password)
	req.Header.Set("Content-Type", c.contentType)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c client) execute(method string, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	req, err := c.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp, err
}

func (c client) Get(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodGet, endpoint, payload)
}

func (c client) Post(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPost, endpoint, payload)
}

func (c client) Put(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPut, endpoint, payload)
}

func (c client) Delete(endpoint string) ([]byte, *http.Response, error) {
	return c.execute(http.MethodDelete, endpoint, nil)
}