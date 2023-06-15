package legacy

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/vimsucks/go-nexus-client/nexus3/pkg/tools"
	"github.com/vimsucks/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryNugetProxy(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryNugetProxy("test-legacy-nuget-repo-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {

		err := service.Delete(createdRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestLegacyRepositoryNugetProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Format: repository.RepositoryFormatNuget,
		Name:   name,
		Online: true,
		Type:   repository.RepositoryTypeProxy,

		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		HTTPClient: &repository.HTTPClient{
			Connection: &repository.HTTPClientConnection{
				Timeout: tools.GetIntPointer(20),
			},
		},
		NegativeCache: &repository.NegativeCache{},
		NugetProxy: &repository.NugetProxy{
			QueryCacheItemMaxAge: 1,
			NugetVersion:         repository.NugetVersion3,
		},
		Proxy: &repository.Proxy{
			RemoteURL: "https://www.nuget.org/api/v2/",
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
	}
}
