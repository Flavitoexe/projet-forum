package router

import (
	"bluedit/forum-frontend/controllers_client"
	"bluedit/forum-frontend/helper"
	"bluedit/forum-frontend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func ClientRoutes(router *mux.Router) {
	RegisterAssetsRoutes(router)
	RegisterScriptsRoutes(router)
	HtmlRoutes(router)
}

func RegisterAssetsRoutes(r *mux.Router) {
	assetsDir := helper.ResolveDir("assets")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))
}

func RegisterScriptsRoutes(r *mux.Router) {
	scriptsDir := helper.ResolveDir("scripts")
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir(scriptsDir))))
}

func HtmlRoutes(router *mux.Router) {

	router.HandleFunc("/", controllers_client.DisplayIndex)
	router.HandleFunc("/thread/{id}", controllers_client.DisplayThread)
	router.HandleFunc("/category/{id}", controllers_client.DisplayCategory)
	router.HandleFunc("/populars", controllers_client.DisplayPopulars)
	router.HandleFunc("/news", controllers_client.DisplayNews)
	router.HandleFunc("/news", controllers_client.DisplayNews)
	router.HandleFunc("/categories", controllers_client.DisplayCategories)

	router.Handle("/message/{id}", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.PostMessage))).Methods("POST")
	router.Handle("/thread", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.PostThread))).Methods("POST")

	router.Handle("/delete/thread/{idCategory}&{idThread}", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.DeleteThread))).Methods("POST")
	router.Handle("/update/thread/{idThread}", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.UpdateThread))).Methods("POST")

	router.Handle("/delete/message/{idThread}&{idMessage}", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.DeleteMessage))).Methods("POST")
	router.Handle("/update/message/{idThread}&{idMessage}", middleware.AuthMiddleware(http.HandlerFunc(controllers_client.UpdateMessage))).Methods("POST")

	router.HandleFunc("/error/{id}", controllers_client.DisplayError)

}
