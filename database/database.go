package database

import (
	"github.com/MizouziE/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("developer:developer@/developer?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Article{}, &models.Region{}, &models.Source{}, &models.Keyword{})

}
