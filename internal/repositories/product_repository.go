package repositories

import (
	"myshop/internal/models"

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

// FindAllWithFilters возвращает список всех товаров с учетом параметров фильтрации и пагинации.
func (r *ProductRepository) FindAllWithFilters(category string, minPrice, maxPrice float64, page, pageSize int) ([]models.Product, error) {
	var products []models.Product
	query := r.DB

	// Применяем фильтрацию по категории, если она указана
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Применяем фильтрацию по цене, если указаны минимальная и/или максимальная цены
	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	// Проверка на корректность параметров пагинации
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Значение по умолчанию
	}

	// Выполняем запрос с учетом параметров пагинации
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	// Получаем список товаров
	result := query.Find(&products)
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
