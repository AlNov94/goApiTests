package steps

import (
	"goApiTests/internal/entity"
	"goApiTests/internal/repository"
)

type UserRepositorySteps struct{}

func (UserRepositorySteps UserRepositorySteps) FindUserById(id int) entity.User {
	return Step("Find user in database by id",
		func() entity.User {
			return repository.GetUserRepository().FindUserById(id)
		},
		"id", id)
}
