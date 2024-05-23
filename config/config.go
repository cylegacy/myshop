package config

import (
	"log"
	"myshop/mocks"
	"myshop/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, got error %v", err)
	}

	// Автоматическая миграция для создания таблицы products
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("failed to migrate database, got error %v", err)
	}

	// Загрузка моковских данных
	mocks.LoadMockData(DB)

	return DB
}
