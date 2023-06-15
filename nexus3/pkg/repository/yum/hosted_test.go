package yum

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/vimsucks/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestYumHostedRepository(name string) repository.YumHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	yumDeployPolicy := repository.YumDeployPolicyPermissive
	return repository.YumHostedRepository{
		Name:   name,
		Online: true,

		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
		Component: &repository.Component{
			ProprietaryComponents: true,
		},
		Yum: repository.Yum{
			RepodataDepth: 10,
			DeployPolicy:  &yumDeployPolicy,
		},
	}
}

func TestYumHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestYumHostedRepository("test-yum-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Component, generatedRepo.Component)
	assert.Equal(t, repo.Yum, generatedRepo.Yum)

	newYumDeployPolicy := repository.YumDeployPolicyStrict
	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Yum.DeployPolicy = &newYumDeployPolicy
	updatedRepo.Yum.RepodataDepth = 5

	err = service.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Yum, generatedRepo.Yum)

	service.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}
