package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/inconshreveable/log15"
	"goApiTests/internal/framework/property"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connectionManager struct {
	dsn string
}

var instance connectionManager

func GetConnectionManager() connectionManager {
	return instance
}

func init() {
	propertyManager := property.GetPropertyManagerInstance()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		propertyManager.GetProperty("db.host"), propertyManager.GetProperty("db.user"), propertyManager.GetProperty("db.password"),
		propertyManager.GetProperty("db.name"), propertyManager.GetProperty("db.port"), propertyManager.GetProperty("db.sslmode"),
		propertyManager.GetProperty("db.timezone"))
	instance = connectionManager{dsn: dsn}
	_, err := gorm.Open(postgres.Open(instance.dsn), &gorm.Config{})
	if err != nil {
		log15.Debug(err.Error())
		panic(err)
	}
	log15.Debug("Connected to database")
}

func (cm connectionManager) GetConnection() *gorm.DB {
	sqlDB, _ := sql.Open("pgx", cm.dsn)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log15.Debug(err.Error())
		panic(err)
	}
	return gormDB.Debug()
}
