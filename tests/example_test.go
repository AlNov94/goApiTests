package tests

import (
	"github.com/stretchr/testify/assert"
	"goApiTests/goApiTests/internal/api"
	"goApiTests/goApiTests/internal/dto"
	"goApiTests/goApiTests/internal/repository"
	"testing"
)

func TestDatabase(t *testing.T) {
	assert := assert.New(t)
	u := repository.GetUserRepository().FindUserByID(1)
	assert.Equal(u.FirstName, "Joe")
}

func TestApi(t *testing.T) {
	assert := assert.New(t)
	var u dto.UserResponse
	api.GetMockService().GetMockUser("1", &u)
	assert.Equal(u.FirstName, "Joe")
}
