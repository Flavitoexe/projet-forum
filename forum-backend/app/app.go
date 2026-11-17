package app

import (
	"bluedit/forum-backend/config"
	"bluedit/forum-backend/controllers"
	repository "bluedit/forum-backend/repositories"
	"bluedit/forum-backend/routers"
	"bluedit/forum-backend/services_api"
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Db        *sql.DB
	RouterApi *mux.Router
}

func InitApp() *App {
	config.LoadEnv()

	db := config.InitDb()

	// Init des repo
	threadRepository := repository.InitThreadRepository(db)
	messageRepository := repository.InitMessageRepository(db)
	categoryRepository := repository.InitCategoryRepository(db)
	userRepository := repository.InitUserRepository(db)
	messageLikeRepository := repository.InitMessageLikeRepositoryRepository(db)
	threadLikeRepository := repository.InitThreadLikeRepositoryRepository(db)

	// Init des services
	threadService := services_api.InitThreadService(threadRepository)
	messageService := services_api.InitMessageService(messageRepository)
	categoryService := services_api.InitCategoryService(categoryRepository)
	userService := services_api.InitUserService(userRepository)
	messageLikeService := services_api.InitMessageLikeService(messageLikeRepository)
	threadLikeService := services_api.InitThreadLikeServices(threadLikeRepository)

	authService := services_api.InitAuthService(userRepository)

	// Init des controllers
	threadController := controllers.InitThreadControllers(threadService)
	messageController := controllers.InitMessageControllers(messageService)
	categoryController := controllers.InitCategoryControllers(categoryService)
	userController := controllers.InitUserControllers(userService)
	messageLikeController := controllers.InitMessageLikeControllers(messageLikeService)
	threadLikeController := controllers.InitThreadLikeControllers(threadLikeService)
	authController := controllers.AuthController(authService)

	routerApi := mux.NewRouter()

	routers.ApiRoutes(routerApi,
		threadController,
		messageController,
		categoryController,
		userController,
		messageLikeController,
		threadLikeController,
		authController,
	)

	return &App{
		Db:        db,
		RouterApi: routerApi,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
	}
}
