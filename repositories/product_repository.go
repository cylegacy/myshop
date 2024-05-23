package repositories

import (
	"myshop/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	result := r.DB.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	result := r.DB.First(&product, id)
	return product, result.Error
}

func (r *ProductRepository) Create(product models.Product) (models.Product, error) {
	result := r.DB.Create(&product)
	return product, result.Error
}

func (r *ProductRepository) Update(product models.Product) (models.Product, error) {
	result := r.DB.Save(&product)
	return product, result.Error
}

func (r *ProductRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Product{}, id)
	return result.Error
}
