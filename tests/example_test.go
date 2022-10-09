package tests

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/steps"

	"sync"
	"testing"

	"github.com/dailymotion/allure-go"
)

var lock sync.Mutex = sync.Mutex{}
var mockServiceApiSteps = steps.MockServiceApiSteps{}
var userRepositorySteps = steps.UserRepositorySteps{}
var assertionSteps = steps.AssertionSteps{}

func TestDatabase(t *testing.T) {
	t.Parallel()
	allure.Test(t, "database test", func() {
		user := userRepositorySteps.FindUserById(1)
		assertionSteps.CheckUserFirstName(t, user, "John")
	})
}

func TestApi(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		lock.Lock()
		defer lock.Unlock()
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
	})
}

func TestApi2(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		lock.Lock()
		defer lock.Unlock()
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
	})
}

func TestApi3(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "Joe")
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
	})
}

func TestApi4(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
	})
}
