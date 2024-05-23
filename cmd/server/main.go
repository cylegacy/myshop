package main

import (
	"log"
	"myshop/config"
	"myshop/controllers"
	"myshop/mocks"
	"myshop/repositories"
	"myshop/routers"
	"myshop/services"
	"net/http"
)

func main() {
	config.InitDB()

	db := config.DB
	mocks.LoadMockData(db)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	router := routers.InitRoutes(productController)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
