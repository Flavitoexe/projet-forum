package routers

import (
	"bluedit/forum-backend/controllers"
	"bluedit/forum-frontend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiRoutes(router *mux.Router,
	threadController *controllers.ThreadControllers,
	messageController *controllers.MessageControllers,
	categoryController *controllers.CategoryControllers,
	userController *controllers.UserControllers,
	messageLikesController *controllers.MessageLikeControllers,
	threadLikesController *controllers.ThreadLikeControllers,
	authController *controllers.AuthControllers,
) {
	ThreadsRoutes(router, threadController)
	MessagesRoutes(router, messageController)
	CategoriesRoutes(router, categoryController)
	UsersRoutes(router, userController)
	MessageLikesRoutes(router, messageLikesController)
	ThreadLikesRoutes(router, threadLikesController)
	AuthRoutes(router, authController)
}

func ThreadsRoutes(router *mux.Router, threadController *controllers.ThreadControllers) {
	router.HandleFunc("/api/thread", threadController.Create).Methods("POST")
	router.HandleFunc("/api/threads", threadController.ReadAll).Methods("GET")
	router.HandleFunc("/api/thread/{id}", threadController.ReadById).Methods("GET")
	router.HandleFunc("/api/thread/{id}", threadController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/thread/{id}", threadController.DeleteById).Methods("DELETE")
}

func MessagesRoutes(router *mux.Router, messageController *controllers.MessageControllers) {
	router.HandleFunc("/api/message", messageController.Create).Methods("POST")
	router.HandleFunc("/api/messages", messageController.ReadAll).Methods("GET")
	router.HandleFunc("/api/message/{id}", messageController.ReadById).Methods("GET")
	router.HandleFunc("/api/message/{id}", messageController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/message/{id}", messageController.DeleteById).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", messageController.ReadByThread).Methods("GET")
}

func CategoriesRoutes(router *mux.Router, categoryController *controllers.CategoryControllers) {
	router.HandleFunc("/api/category", categoryController.Create).Methods("POST")
	router.HandleFunc("/api/categories", categoryController.ReadAll).Methods("GET")
	router.HandleFunc("/api/category/{id}", categoryController.ReadById).Methods("GET")
	router.HandleFunc("/api/category/{id}", categoryController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/category/{id}", categoryController.DeleteById).Methods("DELETE")
}

func UsersRoutes(router *mux.Router, userController *controllers.UserControllers) {
	router.HandleFunc("/api/user", userController.Create).Methods("POST")
	router.HandleFunc("/api/users", userController.ReadAll).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.ReadById).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/user/{id}", userController.DeleteById).Methods("DELETE")
}

func MessageLikesRoutes(router *mux.Router, messageLikesController *controllers.MessageLikeControllers) {
	router.HandleFunc("/api/message_like", messageLikesController.Create).Methods("POST")
	router.HandleFunc("/api/message_likes", messageLikesController.ReadAll).Methods("GET")
	router.HandleFunc("/api/message_like/{id}", messageLikesController.ReadById).Methods("GET")
	router.HandleFunc("/api/message_like/{id}", messageLikesController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/message_like/{id}", messageLikesController.DeleteById).Methods("DELETE")
}

func ThreadLikesRoutes(router *mux.Router, threadLikesController *controllers.ThreadLikeControllers) {
	router.HandleFunc("/api/thread_like", threadLikesController.Create).Methods("POST")
	router.HandleFunc("/api/thread_likes", threadLikesController.ReadAll).Methods("GET")
	router.HandleFunc("/api/thread_like/{id}", threadLikesController.ReadById).Methods("GET")
	router.HandleFunc("/api/thread_like/{id}", threadLikesController.UpdateById).Methods("PUT")
	router.HandleFunc("/api/thread_like/{id}", threadLikesController.DeleteById).Methods("DELETE")
}

func AuthRoutes(router *mux.Router, authController *controllers.AuthControllers) {
	router.HandleFunc("/api/login", authController.Login).Methods("POST")
	router.HandleFunc("/api/register", authController.Register).Methods("POST")

	router.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(authController.Me))).Methods("GET")
}
