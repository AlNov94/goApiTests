package tests

import (
	"goApiTests/goApiTests/internal/api"
	"goApiTests/goApiTests/internal/dto"
	"goApiTests/goApiTests/internal/repository"
	"testing"
)

func TestDatabase(t *testing.T) {
	u := repository.GetUserRepository().FindUserByID(1)
	if u.FirstName != "Joe" {
		t.Error("Expected Joe but was ", u.FirstName)
	}
}

func TestApi(t *testing.T) {
	var u dto.UserResponse
	api.GetMockService().GetMockUser("1", &u)
	if u.FirstName != "Joe" {
		t.Error("Expected Joe but was ", u.FirstName)
	}
}
