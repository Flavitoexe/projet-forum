package main

import (
	"bluedit/forum-backend/app"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := mux.NewRouter()
	router.PathPrefix("/api/").Handler(app.RouterApi)
	router.Handle("/health", app.RouterApi)
	router.Handle("/me", app.RouterApi)
	router.PathPrefix("/").Handler(app.FrontendRouter)

	log.Printf("Serveur front et API lancé : http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erreur lancement serveur - Erreur\n\t%s", err.Error())
	}

