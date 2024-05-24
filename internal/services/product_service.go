package services

import (
	"myshop/internal/models"
	"myshop/internal/repositories"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repository: repo}
}

// GetAllProducts возвращает список всех товаров с учетом параметров фильтрации и пагинации.
func (s *ProductService) GetAllProducts(category string, minPrice, maxPrice float64, page, pageSize int) ([]models.Product, error) {
	return s.Repository.FindAllWithFilters(category, minPrice, maxPrice, page, pageSize)
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.Repository.FindByID(id)
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	return s.Repository.Create(product)
}

func (s *ProductService) UpdateProduct(product models.Product) (models.Product, error) {
	return s.Repository.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repository.Delete(id)
}
