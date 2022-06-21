package repository

import (
	"goApiTests/internal/entity"
	"goApiTests/internal/framework/postgresql"
)

type userRepository struct{}

var userRepositoryInstance userRepository = userRepository{}

func GetUserRepository() userRepository {
	return userRepositoryInstance
}

func (ur userRepository) FindUserByID(id int) entity.User {
	var user entity.User
	postgresql.GetConnectionManager().GetConnection().Where("id = ?", id).Find(&user)
	postgresql.GetConnectionManager().GetConnection().Model(&user).Association("UserProperties").Find(&user.UserProperties)
	return user
}

func (ur userRepository) CreateUser(user *entity.User) {
	postgresql.GetConnectionManager().GetConnection().Create(user)
}

func (ur userRepository) UpdateUser(user *entity.User) {
	postgresql.GetConnectionManager().GetConnection().Updates(user)
}

func (ur userRepository) DeleteUser(user *entity.User) {
	postgresql.GetConnectionManager().GetConnection().Delete(user)
}
