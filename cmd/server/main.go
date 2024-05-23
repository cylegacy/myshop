package main

import (
	"log"
	"myshop/config"
	"myshop/controllers"
	"myshop/middleware"
	"myshop/repositories"
	"myshop/services"
	"net/http"

	_ "myshop/docs" // swagger files

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/dig"
)

// @title My Shop API
// @version 1.0
// @description This is a sample server for a shop.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	container := buildContainer()
	err := container.Invoke(func(router *mux.Router) {
		// Swagger route
		router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		log.Println("Server is running on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	})

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.InitDB)
	container.Provide(repositories.NewProductRepository)
	container.Provide(services.NewProductService)
	container.Provide(controllers.NewProductController)

	// Обертка функции Login как http.HandlerFunc
	container.Provide(func() http.HandlerFunc {
		return controllers.Login
	})

	container.Provide(func(p *controllers.ProductController, loginHandler http.HandlerFunc) *mux.Router {
		router := mux.NewRouter()

		// Публичные маршруты
		router.HandleFunc("/login", loginHandler).Methods("POST")

		// Защищенные маршруты
		api := router.PathPrefix("/api").Subrouter()
		api.Use(middleware.JwtAuthentication)
		api.HandleFunc("/products", p.GetProducts).Methods("GET")
		api.HandleFunc("/products/{id}", p.GetProduct).Methods("GET")
		api.HandleFunc("/products", p.CreateProduct).Methods("POST")
		api.HandleFunc("/products/{id}", p.UpdateProduct).Methods("PUT")
		api.HandleFunc("/products/{id}", p.DeleteProduct).Methods("DELETE")

		return router
	})

	return container
}
