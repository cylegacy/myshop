// internal/routes/auth_routes.go

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes регистрирует маршруты для контроллера авторизации.
func RegisterAuthRoutes(router *mux.Router, loginHandler http.HandlerFunc) {
	router.HandleFunc("/login", loginHandler).Methods("POST")
}
