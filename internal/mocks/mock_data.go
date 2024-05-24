package mocks

import (
	"myshop/internal/models"

	"gorm.io/gorm"
)

func LoadMockData(db *gorm.DB) {
	products := []models.Product{
		{Name: "Laptop", Category: "Electronics", Price: 1000.00},
		{Name: "Shoes", Category: "Fashion", Price: 50.00},
		{Name: "Watch", Category: "Accessories", Price: 200.00},
	}

	for _, product := range products {
		db.Create(&product)
	}
}
