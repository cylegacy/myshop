// internal/routes/product_routes.go

package routes

import (
	"myshop/internal/controllers"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes регистрирует маршруты для контроллера продуктов.
func RegisterProductRoutes(router *mux.Router, productController *controllers.ProductController) {
	router.HandleFunc("/products", productController.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productController.GetProduct).Methods("GET")
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")
}
