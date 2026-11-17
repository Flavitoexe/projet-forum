package router

import (
	"bluedit/forum-frontend/controllers_client"
	"bluedit/forum-frontend/services_client"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	authService := services_client.InitAuthService()
	authController := controllers_client.InitAuthController(authService)

	router.HandleFunc("/login", controllers_client.DisplayLogin).Methods("GET")
	router.HandleFunc("/login", controllers_client.PostLogin).Methods("POST")
	router.HandleFunc("/logout", controllers_client.LogoutUser).Methods("GET")
	router.HandleFunc("/register", controllers_client.DisplayRegister).Methods("GET")
	router.HandleFunc("/register", authController.PostRegister).Methods("POST")
}
