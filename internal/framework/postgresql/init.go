package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/inconshreveable/log15"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	. "restApiTests/framework/property"
)

var dsn string

func init() {
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		GetProperty("db.host"), GetProperty("db.user"), GetProperty("db.password"),
		GetProperty("db.name"), GetProperty("db.port"), GetProperty("db.sslmode"),
		GetProperty("db.timezone"))
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log15.Debug(err.Error())
		panic(err)
	}
	log15.Debug("Connected to database")
}

func GetConnection() *gorm.DB {
	sqlDB, _ := sql.Open("pgx", dsn)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log15.Debug(err.Error())
		panic(err)
	}
	return gormDB.Debug()
}
