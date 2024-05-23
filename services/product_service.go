package services

import (
	"myshop/models"
	"myshop/repositories"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repository: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repository.FindAll()
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
