package utils

import (
	"log"

	"github.com/Wanderer0074348/GoWriteIt/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("file:blog.db?cache=shared&_fk=1"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&models.Post{})

	DB = database
}
