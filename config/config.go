package config

import (
	"log"
	"myshop/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, got error %v", err)
	}

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("failed to migrate database, got error %v", err)
	}
}
