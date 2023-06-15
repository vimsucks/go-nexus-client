package security

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/vimsucks/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestSecurityRoleRead(t *testing.T) {
	service := getTestService()
	role, err := service.Role.Get("nx-admin")

	assert.Nil(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, role.ID, "nx-admin")
	assert.Equal(t, role.Name, "nx-admin")
	assert.Equal(t, 1, len(role.Privileges))
	assert.Equal(t, "nx-all", role.Privileges[0])
	assert.Equal(t, 0, len(role.Roles))

}

func TestSecurityRoleCreateReadUpdateDelete(t *testing.T) {
	service := getTestService()
	roleID := "test-role-" + strconv.Itoa(rand.Intn(1024))
	testRole := testRole(roleID, "test-role-name", "test-role-description", []string{"nx-all"}, []string{"nx-admin"})

	// Create
	err := service.Role.Create(*testRole)
	assert.Nil(t, err)

	if err != nil {
		// Read
		createdRole, err := service.Role.Get(testRole.ID)
		assert.Nil(t, err)
		assert.NotNil(t, createdRole)
		assert.Equal(t, testRole.ID, createdRole.ID)
		assert.Equal(t, testRole.Name, createdRole.Name)
		assert.Equal(t, testRole.Description, createdRole.Description)
		assert.Equal(t, len(testRole.Privileges), len(createdRole.Privileges))
		assert.Equal(t, len(testRole.Roles), len(createdRole.Roles))

		createdRole.Description = "changed"
		createdRole.Name = "changed"
		createdRole.Privileges = []string{"nx-repository-view-*-*-*"}
		createdRole.Roles = []string{"nx-anonymous"}

		// Update
		err = service.Role.Update(createdRole.ID, *createdRole)
		assert.Nil(t, err)

		updatedRole, err := service.Role.Get(createdRole.ID)
		assert.Nil(t, err)
		assert.NotNil(t, updatedRole)
		assert.Equal(t, "changed", updatedRole.Description)
		assert.Equal(t, "changed", updatedRole.Name)
		assert.Equal(t, []string{"nx-repository-view-*-*-*"}, updatedRole.Privileges)
		assert.Equal(t, []string{"nx-anonymous"}, updatedRole.Roles)

		// Delete
		err = service.Role.Delete(createdRole.ID)
		assert.Nil(t, err)

		role, err := service.Role.Get(createdRole.ID)
		assert.Nil(t, err)
		assert.Nil(t, role)
	}
}

func testRole(id, name, description string, privileges []string, roles []string) *security.Role {
	return &security.Role{
		ID:          id,
		Name:        name,
		Description: description,
		Privileges:  privileges,
		Roles:       roles,
	}
}
