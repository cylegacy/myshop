package controllers

import (
	"encoding/json"
	"myshop/internal/models"
	"myshop/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

// @Summary Get a list of all products with filtering and pagination parameters
// @Description Get a list of all products with filtering and pagination parameters.
// @Tags Products
// @Accept json
// @Produce json
// @Param category query string false "Product category"
// @Param minPrice query number false "Minimum product price"
// @Param maxPrice query number false "Maximum product price"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} models.Product
// @Router /api/products [get]
func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры фильтрации из запроса
	params := r.URL.Query()
	category := params.Get("category")
	minPrice, _ := strconv.ParseFloat(params.Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(params.Get("maxPrice"), 64)
	page, _ := strconv.Atoi(params.Get("page"))
	pageSize, _ := strconv.Atoi(params.Get("pageSize"))

	// Передаем параметры в сервис для получения списка товаров
	products, err := c.Service.GetAllProducts(category, minPrice, maxPrice, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем список товаров в ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [get]
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.Service.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Product to create"
// @Success 201 {object} models.Product
// @Router /api/products [post]
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdProduct, err := c.Service.CreateProduct(product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduct)
}

// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [put]
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	product.ID = uint(id)
	updatedProduct, err := c.Service.UpdateProduct(product)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 204
// @Router /api/products/{id} [delete]
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = c.Service.DeleteProduct(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
