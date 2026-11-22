package main

import (
	"bluedit/forum-backend/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Serveur API lancé : http://localhost:%s", port)
	apiServeErr := http.ListenAndServe(":"+port, app.RouterApi)
	if apiServeErr != nil {
		log.Fatalf("Erreur lancement serveur API - Erreur\n\t%s", apiServeErr.Error())
	}
}
