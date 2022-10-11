package steps

import (
	"goApiTests/internal/entity"
	"goApiTests/internal/repository"
	baseSteps "goApiTests/internal/steps/base"
)

//database steps example

type UserRepositorySteps struct{}

func (UserRepositorySteps UserRepositorySteps) FindUserById(id int) entity.User {
	return baseSteps.Step("Find user in database by id",
		func() entity.User {
			return repository.GetUserRepository().FindUserById(id)
		},
		baseSteps.Parameter{Name: "id", Value: id})
}
