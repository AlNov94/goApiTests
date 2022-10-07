package tests

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/repository"
	"goApiTests/internal/steps"
	"sync"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

var lock sync.Mutex = sync.Mutex{}

func (suite *ApiTestSuite) TestDatabase() {
	allure.Test(suite.T(), "database test", func() {
		lock.Lock()
		defer lock.Unlock()
		assert := assert.New(suite.T())
		user := repository.GetUserRepository().FindUserById(1)
		assert.Equal(user.FirstName, "John")
	})
}

func (suite *ApiTestSuite) TestApi() {
	allure.Test(suite.T(), "api test", func() {
		lock.Lock()
		defer lock.Unlock()
		assert := assert.New(suite.T())
		var userResponse dto.UserResponse
		steps.GetMockUserStep("1", &userResponse)
		assert.Equal(userResponse.FirstName, "Joe")
	})
}
