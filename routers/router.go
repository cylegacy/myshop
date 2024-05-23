package routers

import (
	"myshop/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes(controller *controllers.ProductController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controller.GetProduct).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controller.DeleteProduct).Methods("DELETE")
	return router
}
