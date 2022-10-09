package steps

import (
	"fmt"
	"goApiTests/internal/entity"
	"goApiTests/internal/repository"

	"github.com/dailymotion/allure-go"
)

type UserRepositorySteps struct{}

func (UserRepositorySteps UserRepositorySteps) FindUserById(id int) entity.User {
	var result entity.User
	allure.Step(fmt.Sprintf("Find user in database by id %d", id), func() {
		result = repository.GetUserRepository().FindUserById(id)
	})
	return result
}
