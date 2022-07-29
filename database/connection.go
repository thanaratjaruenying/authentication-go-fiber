package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jwt/model"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(postgres.Open("postgresql://test:1234@localhost/test"), &gorm.Config{})
	if err != nil {
		panic("cannot connect to postgres database")
	}

	migrateAuto := connection.AutoMigrate(&model.User{})
	if migrateAuto != nil {
		panic("cannot run migration")
	}

	DB = connection
}
