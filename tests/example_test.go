package tests

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/steps"

	//"sync"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/inconshreveable/log15"
)

// var lock sync.Mutex = sync.Mutex{}
var mockServiceApiSteps = steps.MockServiceApiSteps{}
var userRepositorySteps = steps.UserRepositorySteps{}
var assertionSteps = steps.AssertionSteps{}

func TestDatabase(t *testing.T) {
	t.Parallel()
	allure.Test(t, "database test", func() {
		//lock.Lock()
		//defer lock.Unlock()
		log15.Info("1s")
		user := userRepositorySteps.FindUserById(1)
		assertionSteps.CheckUserFirstName(t, user, "John")
		log15.Info("1f")
	})
}

func TestApi(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		//lock.Lock()
		//defer lock.Unlock()
		log15.Info("2s")
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
		log15.Info("2f")
	})
}

func TestApi2(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		//lock.Lock()
		//defer lock.Unlock()
		log15.Info("3s")
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
		log15.Info("3f")
	})
}

func TestApi3(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		//lock.Lock()
		//defer lock.Unlock()
		log15.Info("4s")
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
		log15.Info("4f")
	})
}

func TestApi4(t *testing.T) {
	t.Parallel()
	allure.Test(t, "api test", func() {
		//lock.Lock()
		//defer lock.Unlock()
		log15.Info("5s")
		var userResponse dto.UserResponse
		mockServiceApiSteps.GetMockUserStep("1", &userResponse)
		assertionSteps.CheckUserResponseFirstName(t, userResponse, "John")
		log15.Info("5f")
	})
}
