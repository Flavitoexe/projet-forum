package app

import (
	"bluedit/forum-backend/config"
	"bluedit/forum-frontend/helper"
	"bluedit/forum-frontend/router"
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Db           *sql.DB
	RouterClient *mux.Router
}

func InitApp() *App {
	config.LoadEnv()

	helper.Load()

	routerClient := mux.NewRouter()

	router.ClientRoutes(routerClient)
	router.AuthRoutes(routerClient)

	return &App{
		RouterClient: routerClient,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
	}
}
