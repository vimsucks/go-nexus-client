package bower

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/vimsucks/go-nexus-client/nexus3/pkg/tools"
	"github.com/vimsucks/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestBowerProxyRepository(name string) repository.BowerProxyRepository {
	return repository.BowerProxyRepository{
		Name:   name,
		Online: true,
		HTTPClient: repository.HTTPClient{
			Blocked:   true,
			AutoBlock: true,
			Connection: &repository.HTTPClientConnection{
				Timeout:       tools.GetIntPointer(20),
				UseTrustStore: tools.GetBoolPointer(true),
			},
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		Proxy: repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "https://archive.ubuntu.com/ubuntu/",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},

		Bower: repository.Bower{
			RewritePackageUrls: true,
		},
	}
}

func TestBowerProxyRepository(t *testing.T) {
	service := getTestService()
	repo := getTestBowerProxyRepository("test-bower-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Proxy.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Connection.Timeout, generatedRepo.HTTPClient.Connection.Timeout)
	assert.Equal(t, repo.HTTPClient.Connection.UseTrustStore, generatedRepo.HTTPClient.Connection.UseTrustStore)
	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Bower, generatedRepo.Bower)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.RewritePackageUrls = false

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Bower, generatedRepo.Bower)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
