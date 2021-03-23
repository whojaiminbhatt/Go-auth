package database

import (
	"github.com/go-projects/go-auth/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:password@/goauth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}