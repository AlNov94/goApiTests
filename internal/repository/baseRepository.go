package repository

import (
	"goApiTests/internal/framework/postgresql"

	"gorm.io/gorm"
)

func getConnection() *gorm.DB {
	return postgresql.GetConnectionManager().GetConnection()
}
