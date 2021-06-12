package router

import (
	"server/controllers"
	"server/controllers/auth"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	// Setting up endpoints
	router.HandleFunc("/api/home", controllers.Home).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/gettoken", auth.GetToken).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", auth.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/register", auth.RegisterHandler).Methods("POST", "OPTIONS")
	return router
}