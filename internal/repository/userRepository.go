package repository

import (
	"goApiTests/internal/entity"
)

type userRepository struct{}

var userRepositoryInstance userRepository = userRepository{}

func GetUserRepository() userRepository {
	return userRepositoryInstance
}

func (ur userRepository) FindUserById(id int) entity.User {
	var user entity.User
	getConnection().Where("id = ?", id).Find(&user)
	getConnection().Model(&user).Association("UserProperties").Find(&user.UserProperties)
	return user
}

func (ur userRepository) CreateUser(user *entity.User) {
	getConnection().Create(user)
}

func (ur userRepository) UpdateUser(user *entity.User) {
	getConnection().Updates(user)
}

func (ur userRepository) DeleteUser(user *entity.User) {
	getConnection().Delete(user)
}
